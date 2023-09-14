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
	"odonto/internal/domain/procedure"
)

// ===================================================================================
// Create Procedure
// ===================================================================================

type createProcedureRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	MID         string  `json:"_mid"` // message identifier
	Version     string  `json:"_v"`   // version
}

type createProcedureResponse struct {
	MID string `json:"_mid"`
}

func decodeCreateProcedureRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(createProcedureRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func makeCreateProcedureEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*createProcedureRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			service := appl.NewProcedureManager(ctx, tx)
			//
			e := &procedure.Entity{
				Name:  req.Name,
				Value: req.Value,
			}
			// create
			return nil, service.Create(e)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &createProcedureResponse{
			MID: req.MID,
		}, nil
	}
}

func CreateProcedureHandler() http.Handler {
	return httptransport.NewServer(
		makeCreateProcedureEndPoint(),
		decodeCreateProcedureRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Find Procedure
// ===================================================================================

type findProcedureRequest struct {
	Pid     string `json:"-"`
	MID     string `json:"_mid"` // message identifier
	Version string `json:"_v"`   // version
}

type findProcedureResponse struct {
	Pid         string  `json:"pid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	MID         string  `json:"_mid"`
}

func decodeFindProcedureRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &findProcedureRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.Pid = vars["pid"]
	return dto, nil
}

func makeFindProcedureEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*findProcedureRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewProcedureManager(ctx, tx)
			return service.Find(&procedure.Identity{
				Pid: req.Pid,
			})
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		result1 := result.(*procedure.Entity)
		// return data
		return &findProcedureResponse{
			Pid:   result1.ID.Pid,
			Name:  result1.Name,
			Value: result1.Value,
			MID:   req.MID,
		}, nil
	}
}

func FindProcedureHandler() http.Handler {
	return httptransport.NewServer(
		makeFindProcedureEndPoint(),
		decodeFindProcedureRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// List Procedure
// ===================================================================================

type listProcedureRequest struct {
	// default
	Text    string `json:"-"`
	Page    int    `json:"-"`
	Limit   int    `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type listProcedureResponse struct {
	Entities []findProcedureResponse `json:"entities"`
	MID      string                  `json:"_mid"`
}

func decodeListProcedureRequest(ctx context.Context, r *http.Request) (interface{}, error) {
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
	dto := &listProcedureRequest{
		Text:    r.URL.Query().Get("text"),
		MID:     r.URL.Query().Get("_mid"),
		Page:    pag,
		Limit:   lim,
		Version: r.URL.Query().Get("_v"),
	}
	return dto, nil
}

func makeListProcedureEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*listProcedureRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		result, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewProcedureManager(ctx, tx)
			// TODO: Remover identity de list
			list, err := service.List(procedure.Identity{}, req.Text, req.Page, req.Limit)
			if err != nil {
				return nil, err
			}
			// result
			entities := make([]findProcedureResponse, 0)
			for _, v := range list {
				entity := findProcedureResponse{
					Pid:   v.ID.Pid,
					Name:  v.Name,
					Value: v.Value,
					MID:   req.MID,
				}
				entities = append(entities, entity)
			}
			return entities, nil
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		entities := result.([]findProcedureResponse)
		// return
		return &listProcedureResponse{
			Entities: entities,
			MID:      req.MID,
		}, nil
	}
}

func ListProcedureHandler() http.Handler {
	return httptransport.NewServer(
		makeListProcedureEndPoint(),
		decodeListProcedureRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Update Procedure
// ===================================================================================

type updateProcedureRequest struct {
	Pid         string  `json:"pid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	MID         string  `json:"_mid"` // message identifier
	Version     string  `json:"_v"`   // version
}

type updateProcedureResponse struct {
	MID string `json:"_mid"`
}

func decodeUpdateProcedureRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	dto := new(updateProcedureRequest)
	decoder := json.NewDecoder(r.Body)
	// get from url params
	vars := mux.Vars(r)
	// decode from body
	err := decoder.Decode(dto)
	if err != nil {
		return nil, err
	}
	// dto pk
	dto.Pid = vars["pid"]
	return dto, nil
}

func makeUpdateProcedureEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*updateProcedureRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewProcedureManager(ctx, tx)
			identity := procedure.Identity{
				Pid: req.Pid,
			}
			e := &procedure.Entity{
				Name:  req.Name,
				Value: req.Value,
				ID:    identity,
			}
			// update
			return nil, service.Update(e)
		})
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &updateProcedureResponse{
			MID: req.MID,
		}, nil
	}
}

func UpdateProcedureHandler() http.Handler {
	return httptransport.NewServer(
		makeUpdateProcedureEndPoint(),
		decodeUpdateProcedureRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}

// ===================================================================================
// Remove Procedure
// ===================================================================================

type removeProcedureRequest struct {
	Pid     string `json:"-"`
	MID     string `json:"-"`
	Version string `json:"-"`
}

type removeProcedureResponse struct {
	MID string `json:"_mid"`
}

func decodeRemoveProcedureRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//vars
	vars := mux.Vars(r)
	dto := &removeProcedureRequest{
		MID:     r.URL.Query().Get("_mid"),
		Version: r.URL.Query().Get("_v"),
	}
	// dto pk
	dto.Pid = vars["pid"]
	return dto, nil
}

func makeRemoveProcedureEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// retrieve request data
		req, ok := request.(*removeProcedureRequest)
		if !ok {
			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
		}
		// process data
		_, err := appl.TxRun(func(tx *sql.Tx) (interface{}, error) {
			// process data
			service := appl.NewProcedureManager(ctx, tx)
			return nil, service.Remove(&procedure.Identity{
				Pid: req.Pid,
			})
		})
		// success
		if err != nil {
			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
		}
		//return data
		return &removeProcedureResponse{
			MID: req.MID,
		}, nil
	}
}

func RemoveProcedureHandler() http.Handler {
	return httptransport.NewServer(
		makeRemoveProcedureEndPoint(),
		decodeRemoveProcedureRequest,
		encodeResponse,
		httptransport.ServerBefore(httpToContext()),
		httptransport.ServerErrorEncoder(errorEncoder()),
	)
}
