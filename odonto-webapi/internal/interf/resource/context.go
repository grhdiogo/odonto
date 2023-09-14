package resource

import (
	"context"
	"net/http"
	"strings"

	httptransporter "github.com/go-kit/kit/transport/http"
)

type appContextKey int8

const (
	appTokenCtxKey appContextKey = iota
	appNameCtxKey
	appVersionCtxKey
	jwtTokenCtxKey
	tenantCtxKey
)

// extracts context variables
//
// HEADERS = {
//   "x-app-name": String,       // name application (required)
//   "x-app-version" : String,   // version application (required)
//   "x-app-token" : String,     // app token for access (required)
//   "x-app-tenant" : String,    // app tenant  (required)
// }
func httpToContext() httptransporter.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		// get values
		xAppToken := r.Header.Get("x-app-token")
		xTenant := r.Header.Get("x-app-tenant")
		xAppName := r.Header.Get("x-app-name")
		xAppVer := r.Header.Get("x-app-version")
		bearerToken := r.Header.Get("authorization")
		bearerToken = strings.Replace(bearerToken, "Bearer", "", 1)
		jwtToken := strings.TrimSpace(bearerToken)
		// set values to context
		newCtx := context.WithValue(ctx, appTokenCtxKey, xAppToken)
		newCtx = context.WithValue(newCtx, appNameCtxKey, xAppName)
		newCtx = context.WithValue(newCtx, appVersionCtxKey, xAppVer)
		newCtx = context.WithValue(newCtx, jwtTokenCtxKey, jwtToken)
		newCtx = context.WithValue(newCtx, tenantCtxKey, xTenant)
		// context
		return newCtx
	}
}