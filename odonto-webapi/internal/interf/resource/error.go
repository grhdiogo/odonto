package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"odonto/internal/infra/fails"

	httptransport "github.com/go-kit/kit/transport/http"
)

// =====================================================================
// Error response
// =====================================================================

type ErrorResponse struct {
	Status int         `json:"status"`
	Code   int         `json:"code"`
	Msg    string      `json:"message"`
	Errs   fails.Stack `json:"errorStack"`
	MID    string      `json:"_mid"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, mid: %s",
		e.Code, e.Msg, e.MID,
	)
}

// =====================================================================
// Error adapter for gokit
// =====================================================================
func errorEncoder() httptransport.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		// check if http error type
		rErr, ok := err.(*ErrorResponse)
		if !ok {
			rErr = &ErrorResponse{
				Status: 500,
				Code:   0,
				Msg:    err.Error(),
				MID:    "na",
			}
		}
		// write status
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(rErr.Status)
		// encode and write error response
		encoder := json.NewEncoder(w)
		encoder.Encode(rErr)
	}
}

// =====================================================================
// Error response creator
// =====================================================================

func CreateHttpErrorResponse(mid string, status, code int, msg string, err error) *ErrorResponse {
	failsStack, ok := err.(fails.Stack)
	if !ok {
		failsStack = fails.ProcessErrStack(err)
	}
	return &ErrorResponse{
		Status: status,
		Code:   code,
		Msg:    msg,
		Errs:   failsStack,
		MID:    mid,
	}
}

func CreateApplErrorResponse(mid string, status int, err error) *ErrorResponse {
	failsStack, ok := err.(fails.Stack)
	if !ok {
		failsStack = fails.ProcessErrStack(err)
	}
	return &ErrorResponse{
		Status: status,
		Code:   failsStack.Code(),
		Msg:    failsStack.Message(),
		Errs:   failsStack.ErrStack(),
		MID:    mid,
	}
}
