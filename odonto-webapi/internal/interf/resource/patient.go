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
	"odonto/internal/domain/patient"
	"odonto/internal/domain/shared"
)

// ===================================================================================
// Create Patient
// ===================================================================================

type createPatientRequest struct {
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
	MID       string `json:"_mid"` // message identifier
	Version   string `json:"_v"`   // version
}

type createPatientResponse struct {
	MID string `json:"_mid"`
}

func decodeCreatePatientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(createPatientRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func makeCreatePatientEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*createPatientRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		return appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			service := appl.NewPatientManager(ctx, tx)
			// create
			err := service.Create(req.Name, req.Cpf, req.Email, req.Birthdate)
			if err != nil {
				return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
			}

			return &createPatientResponse{
				MID: req.MID,
			}, nil
		})
	}
}

func CreatePatientHandler() http.Handler {
	return httptransport.NewServer(
		makeCreatePatientEndPoint(),
		decodeCreatePatientRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Find Patient
// ===================================================================================

type findPatientRequest struct {
	PersonPid string `json:"-"`
	MID       string `json:"_mid"` // message identifier
	Version   string `json:"_v"`   // version
}

type findPatientResponse struct {
	Pid       string `json:"pid"`
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
	MID       string `json:"_mid"`
}

func decodeFindPatientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &findPatientRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.PersonPid = vars["personPid"]
	return dto, nil
}

func makeFindPatientEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*findPatientRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewPatientManager(ctx, tx)
			return service.Find(&patient.Identity{
				PersonPid: req.PersonPid,
			})
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		result1 := result.(*patient.Entity)
		// return data
		return &findPatientResponse{
			Pid:       result1.ID.PersonPid,
			Name:      result1.Name,
			Cpf:       result1.Cpf,
			Email:     result1.Email,
			Birthdate: result1.Birthdate.Format(shared.DefaultTimeLayout),
			MID:       req.MID,
		}, nil
	}
}

func FindPatientHandler() http.Handler {
	return httptransport.NewServer(
		makeFindPatientEndPoint(),
		decodeFindPatientRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// List Patient
// ===================================================================================

type listPatientRequest struct {
	// default
	Text    string `json:"-"`
	Page    int    `json:"-"`
	Limit   int    `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type listPatientResponse struct {
	Entities []findPatientResponse `json:"entities"`
	Total    int                   `json:"total"`
	MID      string                `json:"_mid"`
}

func decodeListPatientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
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
	dto := &listPatientRequest{
		Text:    r.URL.Query().Get("text"),
		MID:     r.URL.Query().Get("_mid"),
		Page:    pag,
		Limit:   lim,
		Version: r.URL.Query().Get("_v"),
	}
	return dto, nil
}

func makeListPatientEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*listPatientRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}

		return appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewPatientManager(ctx, tx)
			list, count, err := service.List(req.Text, req.Page, req.Limit)
			if err != nil {
				return nil, err
			}
			// result
			entities := make([]findPatientResponse, 0)
			for _, v := range list {
				entity := findPatientResponse{
					Pid:       v.ID.PersonPid,
					Name:      v.Name,
					Cpf:       v.Cpf,
					Email:     v.Email,
					Birthdate: v.Birthdate.Format(shared.DefaultTimeLayout),
					MID:       req.MID,
				}
				entities = append(entities, entity)
			}
			return &listPatientResponse{
				Entities: entities,
				Total:    count,
				MID:      req.MID,
			}, nil
		})
	}
}

func ListPatientHandler() http.Handler {
	return httptransport.NewServer(
		makeListPatientEndPoint(),
		decodeListPatientRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Update Patient
// ===================================================================================

type updatePatientRequest struct {
	Pid     string `json:"-"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	MID     string `json:"_mid"` // message identifier
	Version string `json:"_v"`   // version
}

type updatePatientResponse struct {
	MID string `json:"_mid"`
}

func decodeUpdatePatientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(updatePatientRequest)
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

func makeUpdatePatientEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*updatePatientRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewPatientManager(ctx, tx)
			// update
			return nil, service.Update(req.Pid, req.Name, req.Email)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &updatePatientResponse{
			MID: req.MID,
		}, nil
	}
}

func UpdatePatientHandler() http.Handler {
	return httptransport.NewServer(
		makeUpdatePatientEndPoint(),
		decodeUpdatePatientRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Remove Patient
// ===================================================================================

type removePatientRequest struct {
	PersonPid string `json:"-"`
	MID       string `json:"-"`
	Version   string `json:"-"`
}

type removePatientResponse struct {
	MID string `json:"_mid"`
}

func decodeRemovePatientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &removePatientRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.PersonPid = vars["personPid"]
	return dto, nil
}

func makeRemovePatientEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*removePatientRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewPatientManager(ctx, tx)
			return nil, service.Remove(&patient.Identity{
				PersonPid: req.PersonPid,
			})
		})
		// success
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &removePatientResponse{
			MID: req.MID,
		}, nil
	}
}

func RemovePatientHandler() http.Handler {
	return httptransport.NewServer(
		makeRemovePatientEndPoint(),
		decodeRemovePatientRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}
