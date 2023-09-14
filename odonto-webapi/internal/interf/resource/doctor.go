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
	"odonto/internal/domain/doctor"
	"odonto/internal/domain/shared"
)

// ===================================================================================
// Create Doctor
// ===================================================================================

type createDoctorRequest struct {
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
	MID       string `json:"_mid"` // message identifier
	Version   string `json:"_v"`   // version
}

type createDoctorResponse struct {
	MID string `json:"_mid"`
}

func decodeCreateDoctorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(createDoctorRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func makeCreateDoctorEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*createDoctorRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		return appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			service := appl.NewDoctorManager(ctx, tx)
			// create
			err := service.Create(req.Name, req.Cpf, req.Email, req.Birthdate)
			if err != nil {
				return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
			}

			return &createDoctorResponse{
				MID: req.MID,
			}, nil
		})
	}
}

func CreateDoctorHandler() http.Handler {
	return httptransport.NewServer(
		makeCreateDoctorEndPoint(),
		decodeCreateDoctorRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Find Doctor
// ===================================================================================

type findDoctorRequest struct {
	PersonPid string `json:"-"`
	MID       string `json:"_mid"` // message identifier
	Version   string `json:"_v"`   // version
}

type findDoctorResponse struct {
	Pid       string `json:"pid"`
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
	MID       string `json:"_mid"`
}

func decodeFindDoctorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &findDoctorRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.PersonPid = vars["personPid"]
	return dto, nil
}

func makeFindDoctorEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*findDoctorRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}

		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewDoctorManager(ctx, tx)
			return service.Find(&doctor.Identity{
				PersonPid: req.PersonPid,
			})
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		result1 := result.(*doctor.Entity)
		// return data
		return &findDoctorResponse{
			Pid:       result1.ID.PersonPid,
			Name:      result1.Name,
			Cpf:       result1.Cpf,
			Email:     result1.Email,
			Birthdate: result1.Birthdate.Format(shared.DefaultTimeLayout),
			MID:       req.MID,
		}, nil
	}
}

func FindDoctorHandler() http.Handler {
	return httptransport.NewServer(
		makeFindDoctorEndPoint(),
		decodeFindDoctorRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// List Doctor
// ===================================================================================

type listDoctorRequest struct {
	// default
	Text    string `json:"-"`
	Page    int    `json:"-"`
	Limit   int    `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type listDoctorResponse struct {
	Entities []findDoctorResponse `json:"entities"`
	Total    int                  `json:"total"`
	MID      string               `json:"_mid"`
}

func decodeListDoctorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
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
	dto := &listDoctorRequest{
		Text:    r.URL.Query().Get("text"),
		MID:     r.URL.Query().Get("_mid"),
		Page:    pag,
		Limit:   lim,
		Version: r.URL.Query().Get("_v"),
	}
	return dto, nil
}

func makeListDoctorEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*listDoctorRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}

		return appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewDoctorManager(ctx, tx)
			list, count, err := service.List(req.Text, req.Page, req.Limit)
			if err != nil {
				return nil, err
			}
			// result
			entities := make([]findDoctorResponse, 0)
			for _, v := range list {
				entity := findDoctorResponse{
					Pid:       v.ID.PersonPid,
					Name:      v.Name,
					Cpf:       v.Cpf,
					Email:     v.Email,
					Birthdate: v.Birthdate.Format(shared.DefaultTimeLayout),
					MID:       req.MID,
				}
				entities = append(entities, entity)
			}
			return &listDoctorResponse{
				Entities: entities,
				Total:    count,
				MID:      req.MID,
			}, nil
		})
	}
}

func ListDoctorHandler() http.Handler {
	return httptransport.NewServer(
		makeListDoctorEndPoint(),
		decodeListDoctorRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Update Doctor
// ===================================================================================

type updateDoctorRequest struct {
	Pid     string `json:"-"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	MID     string `json:"_mid"` // message identifier
	Version string `json:"_v"`   // version
}

type updateDoctorResponse struct {
	MID string `json:"_mid"`
}

func decodeUpdateDoctorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(updateDoctorRequest)
	decoder := json.NewDecoder(r.Body)
	// get from url params
	vars := mux.Vars(r)
	// decode from body
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	// dto pk
	dto.Pid = vars["personPid"]
	return dto, nil
}

func makeUpdateDoctorEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*updateDoctorRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewDoctorManager(ctx, tx)
			// update
			return nil, service.Update(req.Pid, req.Name, req.Email)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &updateDoctorResponse{
			MID: req.MID,
		}, nil
	}
}

func UpdateDoctorHandler() http.Handler {
	return httptransport.NewServer(
		makeUpdateDoctorEndPoint(),
		decodeUpdateDoctorRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Remove Doctor
// ===================================================================================

type removeDoctorRequest struct {
	PersonPid string `json:"-"`
	MID       string `json:"-"`
	Version   string `json:"-"`
}

type removeDoctorResponse struct {
	MID string `json:"_mid"`
}

func decodeRemoveDoctorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &removeDoctorRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.PersonPid = vars["personPid"]
	return dto, nil
}

func makeRemoveDoctorEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*removeDoctorRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}

		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewDoctorManager(ctx, tx)
			return nil, service.Remove(&doctor.Identity{
				PersonPid: req.PersonPid,
			})
		})
		// success
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &removeDoctorResponse{
			MID: req.MID,
		}, nil
	}
}

func RemoveDoctorHandler() http.Handler {
	return httptransport.NewServer(
		makeRemoveDoctorEndPoint(),
		decodeRemoveDoctorRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}
