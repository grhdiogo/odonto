package resource

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/go-kit/kit/endpoint"
// 	httptransport "github.com/go-kit/kit/transport/http"

// 	"odonto/internal/appl"
// 	"odonto/internal/domain/authaccess"
// )

// // ===================================================================================
// // Create Authaccess
// // ===================================================================================

// type createAuthaccessRequest struct {
// 	Identifier string `json:"identifier"`
// 	Secret     string `json:"secret"`
// 	MID        string `json:"_mid"` // message identifier
// 	Version    string `json:"_v"`   // version
// }

// type createAuthaccessResponse struct {
// 	JwtToken string `json:"jwtToken"`
// 	MID      string `json:"_mid"`
// }

// func decodeCreateAuthaccessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	dto := new(createAuthaccessRequest)
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(dto)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dto, nil
// }

// func makeCreateAuthaccessEndPoint() endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		// retrieve request data
// 		req, ok := request.(*createAuthaccessRequest)
// 		if !ok {
// 			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
// 		}
// 		// process data
// 		service := appl.NewAuthaccessManager(ctx)
// 		// create
// 		token, err := service.Create(req.Identifier, req.Secret)
// 		if err != nil {
// 			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
// 		}
// 		//return data
// 		return &createAuthaccessResponse{
// 			JwtToken: token,
// 			MID:      req.MID,
// 		}, nil
// 	}
// }

// func CreateAuthaccessHandler() http.Handler {
// 	return httptransport.NewServer(
// 		makeCreateAuthaccessEndPoint(),
// 		decodeCreateAuthaccessRequest,
// 		encodeResponse,
// 		httptransport.ServerBefore(httpToContext()),
// 		httptransport.ServerErrorEncoder(errorEncoder()),
// 	)
// }

// // ===================================================================================
// // Check Token
// // ===================================================================================

// type checkTokenRequest struct {
// 	MID     string `json:"_mid"` // message identifier
// 	Version string `json:"_v"`   // version
// }

// type checkTokenResponse struct {
// 	IsValid bool   `json:"isValid"`
// 	MID     string `json:"_mid"`
// }

// func decodeCheckTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	dto := &checkTokenRequest{
// 		MID: r.URL.Query().Get("_mid"),
// 	}
// 	return dto, nil
// }

// func makeCheckTokenEndPoint() endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		// retrieve request data
// 		req, ok := request.(*checkTokenRequest)
// 		if !ok {
// 			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
// 		}
// 		jwtToken := ctx.Value(jwtTokenCtxKey).(string)
// 		// process data
// 		service := appl.NewAuthaccessManager(ctx)
// 		valid := service.CheckToken(jwtToken)
// 		// return data
// 		return &checkTokenResponse{
// 			IsValid: valid,
// 			MID:     req.MID,
// 		}, nil
// 	}
// }

// func CheckTokenHandler() http.Handler {
// 	return httptransport.NewServer(
// 		makeCheckTokenEndPoint(),
// 		decodeCheckTokenRequest,
// 		encodeResponse,
// 		httptransport.ServerBefore(httpToContext()),
// 		httptransport.ServerErrorEncoder(errorEncoder()),
// 	)
// }

// // ===================================================================================
// // Remove Authaccess
// // ===================================================================================

// type removeAuthaccessRequest struct {
// 	MID     string `json:"-"`
// 	Version string `json:"-"`
// }

// type removeAuthaccessResponse struct {
// 	MID string `json:"_mid"`
// }

// func decodeRemoveAuthaccessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	//vars
// 	dto := &removeAuthaccessRequest{
// 		MID:     r.URL.Query().Get("_mid"),
// 		Version: r.URL.Query().Get("_v"),
// 	}
// 	// dto pk
// 	return dto, nil
// }

// func makeRemoveAuthaccessEndPoint() endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		// retrieve request data
// 		req, ok := request.(*removeAuthaccessRequest)
// 		if !ok {
// 			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
// 		}
// 		// process data
// 		jwtToken := ctx.Value(jwtTokenCtxKey).(string)
// 		ent, err := authaccess.NewEntity(jwtToken)
// 		if err != nil {
// 			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
// 		}
// 		service := appl.NewAuthaccessManager(ctx)
// 		err = service.Remove(&authaccess.Identity{
// 			Token: ent.ID.Token,
// 		})
// 		if err != nil {
// 			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
// 		}
// 		//return data
// 		return &removeAuthaccessResponse{
// 			MID: req.MID,
// 		}, nil
// 	}
// }

// func RemoveAuthaccessHandler() http.Handler {
// 	return httptransport.NewServer(
// 		makeRemoveAuthaccessEndPoint(),
// 		decodeRemoveAuthaccessRequest,
// 		encodeResponse,
// 		httptransport.ServerBefore(httpToContext()),
// 		httptransport.ServerErrorEncoder(errorEncoder()),
// 	)
// }

// // ===================================================================================
// // Verify First access
// // ===================================================================================

// type verifyFirstAccessRequest struct {
// 	MID     string `json:"_mid"` // message identifier
// 	Version string `json:"_v"`   // version
// }

// type verifyFirstAccessResponse struct {
// 	IsFirstAccess bool   `json:"isFirstAccess"`
// 	MID           string `json:"_mid"`
// }

// func decodeVerifyFirstAccessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	dto := new(verifyFirstAccessRequest)
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(dto)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return dto, nil
// }

// func makeVerifyFirstAccessEndPoint() endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		// retrieve request data
// 		req, ok := request.(*verifyFirstAccessRequest)
// 		if !ok {
// 			return nil, CreateHttpErrorResponse("na", http.StatusBadRequest, 1001, "invalid request", nil)
// 		}
// 		jwtToken := ctx.Value(jwtTokenCtxKey).(string)
// 		ent, err := authaccess.NewEntity(jwtToken)
// 		if err != nil {
// 			return nil, CreateApplErrorResponse(req.MID, http.StatusBadRequest, err)
// 		}
// 		// process data
// 		service := appl.NewAuthaccessManager(ctx)
// 		valid := service.IsFirstAccess(ent)
// 		// return data
// 		return &verifyFirstAccessResponse{
// 			IsFirstAccess: valid,
// 			MID:           req.MID,
// 		}, nil
// 	}
// }

// func VerifyFirstAccessHandler() http.Handler {
// 	return httptransport.NewServer(
// 		makeVerifyFirstAccessEndPoint(),
// 		decodeVerifyFirstAccessRequest,
// 		encodeResponse,
// 		httptransport.ServerBefore(httpToContext()),
// 		httptransport.ServerErrorEncoder(errorEncoder()),
// 	)
// }
