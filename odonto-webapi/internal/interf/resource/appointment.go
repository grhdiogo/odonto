package resource

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"odonto/internal/appl"
	"odonto/internal/domain/appointment"
)

type Items struct {
	ID    string
	Name  string
	Value float64
	Tooth int
}

// ===================================================================================
// Create Appointment
// ===================================================================================

type createAppointmentRequest struct {
	Observation string  `json:"observation"`
	DoctorDid   string  `json:"doctorDid"`
	PatientPid  string  `json:"patientPid"`
	Items       []Items `json:"items"`
	MID         string  `json:"_mid"` // message identifier
	Version     string  `json:"_v"`   // version
}

type createAppointmentResponse struct {
	MID string `json:"_mid"`
}

func decodeCreateAppointmentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(createAppointmentRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func makeCreateAppointmentEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*createAppointmentRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			service := appl.NewAppointmentManager(ctx, tx)
			//
			e := &appointment.Entity{
				DoctorDid:   req.DoctorDid,
				PatientPid:  req.PatientPid,
				Observation: req.Observation,
			}
			for i := 0; i < len(req.Items); i++ {
				e.Items = append(e.Items, appointment.Item{
					ProcedurePID:   req.Items[i].ID,
					ProcedureName:  req.Items[i].Name,
					ProcedureValue: req.Items[i].Value,
					Tooth:          req.Items[i].Tooth,
				})
			}
			// create
			return nil, service.Create(e)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &createAppointmentResponse{
			MID: req.MID,
		}, nil
	}
}

func CreateAppointmentHandler() http.Handler {
	return httptransport.NewServer(
		makeCreateAppointmentEndPoint(),
		decodeCreateAppointmentRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Find Appointment
// ===================================================================================

type findAppointmentRequest struct {
	Aid     string `json:"-"`
	MID     string `json:"_mid"` // message identifier
	Version string `json:"_v"`   // version
}

type findAppointmentResponse struct {
	Aid         string  `json:"aid"`
	Status      string  `json:"status"`
	Observation string  `json:"observation"`
	DoctorDid   string  `json:"doctorDid"`
	PatientPid  string  `json:"patientPid"`
	DoctorName  string  `json:"doctorName"`
	PatientName string  `json:"patientName"`
	Date        string  `json:"date"`
	Items       []Items `json:"items"`
	MID         string  `json:"_mid"`
}

func decodeFindAppointmentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &findAppointmentRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.Aid = vars["aid"]
	return dto, nil
}

func makeFindAppointmentEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*findAppointmentRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewAppointmentManager(ctx, tx)
			return service.Find(&appointment.Identity{
				Aid: req.Aid,
			})
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		result1 := result.(*appointment.Entity)
		r := &findAppointmentResponse{
			Aid:        result1.ID.Aid,
			Status:     result1.Status,
			DoctorDid:  result1.DoctorDid,
			PatientPid: result1.PatientPid,
			MID:        req.MID,
		}

		for i := 0; i < len(result1.Items); i++ {
			r.Items = append(r.Items, Items{
				Name:  result1.Items[i].ProcedureName,
				Value: result1.Items[i].ProcedureValue,
				Tooth: result1.Items[i].Tooth,
				ID:    result1.Items[i].ProcedurePID,
			})
		}

		// return data
		return r, nil
	}
}

func FindAppointmentHandler() http.Handler {
	return httptransport.NewServer(
		makeFindAppointmentEndPoint(),
		decodeFindAppointmentRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// List Appointment
// ===================================================================================

type listAppointmentRequest struct {
	// default
	Text    string `json:"-"`
	Page    int    `json:"-"`
	Limit   int    `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type listAppointmentResponse struct {
	Entities []findAppointmentResponse `json:"entities"`
	Total    int                       `json:"total"`
	MID      string                    `json:"_mid"`
}

func decodeListAppointmentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// convert to int
	lim, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		// default value
		lim = 10
	}
	pag, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		// default value
		pag = 1
	}
	dto := &listAppointmentRequest{
		Text:    r.URL.Query().Get("text"),
		MID:     r.URL.Query().Get("_mid"),
		Page:    pag,
		Limit:   lim,
		Version: r.URL.Query().Get("_v"),
	}
	return dto, nil
}

func makeListAppointmentEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*listAppointmentRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewAppointmentManager(ctx, tx)
			// TODO: Remover identity de list
			list, count, err := service.List(appointment.Identity{}, req.Text, req.Page, req.Limit)
			if err != nil {
				return nil, err
			}
			// result
			entities := make([]findAppointmentResponse, 0)
			for _, v := range list {
				entity := findAppointmentResponse{
					Aid:         v.ID.Aid,
					Status:      v.Status,
					DoctorDid:   v.DoctorDid,
					PatientPid:  v.PatientPid,
					DoctorName:  v.DoctorName,
					Observation: v.Observation,
					PatientName: v.PatientName,
					Date:        v.CreatedAt.Format("02/01/2006"),
					MID:         req.MID,
				}
				for i := 0; i < len(v.Items); i++ {
					entity.Items = append(entity.Items, Items{
						Name:  v.Items[i].ProcedureName,
						Value: v.Items[i].ProcedureValue,
						Tooth: v.Items[i].Tooth,
						ID:    v.Items[i].ProcedurePID,
					})
				}
				entities = append(entities, entity)
			}
			return &listAppointmentResponse{
				Entities: entities,
				Total:    count,
				MID:      req.MID,
			}, nil
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		// return
		return result, nil
	}
}

func ListAppointmentHandler() http.Handler {
	return httptransport.NewServer(
		makeListAppointmentEndPoint(),
		decodeListAppointmentRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Update Appointment
// ===================================================================================

type updateAppointmentRequest struct {
	Aid         string  `json:"aid"`
	Status      string  `json:"status"`
	Observation string  `json:"observation"`
	DoctorDid   string  `json:"doctorDid"`
	PatientPid  string  `json:"patientPid"`
	Items       []Items `json:"items"`
	MID         string  `json:"_mid"` // message identifier
	Version     string  `json:"_v"`   // version
}

type updateAppointmentResponse struct {
	MID string `json:"_mid"`
}

func decodeUpdateAppointmentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(updateAppointmentRequest)
	decoder := json.NewDecoder(r.Body)
	// get from url params
	vars := mux.Vars(r)
	// decode from body
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	// dto pk
	dto.Aid = vars["aid"]
	return dto, nil
}

func makeUpdateAppointmentEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*updateAppointmentRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewAppointmentManager(ctx, tx)
			identity := appointment.Identity{
				Aid: req.Aid,
			}
			e := &appointment.Entity{
				Status:      req.Status,
				DoctorDid:   req.DoctorDid,
				PatientPid:  req.PatientPid,
				Observation: req.Observation,
				ID:          identity,
			}
			for i := 0; i < len(req.Items); i++ {
				e.Items = append(e.Items, appointment.Item{
					ProcedureName:  req.Items[i].Name,
					ProcedureValue: req.Items[i].Value,
					Tooth:          req.Items[i].Tooth,
					ProcedurePID:   req.Items[i].ID,
				})
			}
			// update
			return nil, service.Update(e)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &updateAppointmentResponse{
			MID: req.MID,
		}, nil
	}
}

func UpdateAppointmentHandler() http.Handler {
	return httptransport.NewServer(
		makeUpdateAppointmentEndPoint(),
		decodeUpdateAppointmentRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Remove Appointment
// ===================================================================================

type removeAppointmentRequest struct {
	Aid     string `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type removeAppointmentResponse struct {
	MID string `json:"_mid"`
}

func decodeRemoveAppointmentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &removeAppointmentRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.Aid = vars["aid"]
	return dto, nil
}

func makeRemoveAppointmentEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*removeAppointmentRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewAppointmentManager(ctx, tx)
			return nil, service.Remove(&appointment.Identity{
				Aid: req.Aid,
			})
		})
		// success
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &removeAppointmentResponse{
			MID: req.MID,
		}, nil
	}
}

func RemoveAppointmentHandler() http.Handler {
	return httptransport.NewServer(
		makeRemoveAppointmentEndPoint(),
		decodeRemoveAppointmentRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}
