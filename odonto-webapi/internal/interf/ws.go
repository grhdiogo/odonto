package ws

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"odonto/internal/infra/config"
	"odonto/internal/interf/routes"
	"strings"

	"github.com/gorilla/mux"
	"github.com/robbert229/jwt"
)

//==============================================================================
//  Middlewares
//==============================================================================

// TODO: check device key

// check authorization
func checkAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check path
		if strings.Contains(r.URL.Path, "/priv/") {
			// recover authorization
			bearerToken := r.Header.Get("authorization")
			bearerToken = strings.Replace(bearerToken, "Bearer", "", 1)
			jwtToken := strings.TrimSpace(bearerToken)
			// get settings
			settings := config.GetSettings()
			// create a jwt encoder
			jwtSecret := settings.JWTSecret
			algorithm := jwt.HmacSha256(jwtSecret)
			// extract stoken from jwt token
			if err := algorithm.Validate(jwtToken); err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Invalid token"))
				return
			}
			// //decode jwttoken
			// claims, err := algorithm.Decode(jwtToken)
			// if err != nil {
			// 	w.WriteHeader(http.StatusForbidden)
			// 	w.Write([]byte("error on decode token"))
			// 	return
			// }
			// //get token
			// algorithm.Decode(jwtToken)
			// token, err := claims.Get("tkn")
			// if err != nil {
			// 	w.WriteHeader(http.StatusForbidden)
			// 	w.Write([]byte("error on get decoded token"))
			// 	return
			// }
			// //verificar token
			// access := applpatient.NewAccessService(r.Context())
			// err = access.CheckToken(token.(string))
			// if err != nil {
			// 	w.WriteHeader(http.StatusForbidden)
			// 	w.Write([]byte("token not found"))
			// 	return
			// }
		}
		// next router
		next.ServeHTTP(w, r)
	})
}

//==============================================================================
//  Interfaces
//==============================================================================

//WebService an abstraction of web service api
type WebService interface {
	// Init configures routes, prefix, and handlers
	Init()
	//GetRouters retrieve configured routers
	GetRouters() http.Handler
}

//==============================================================================
//  Implementation
//==============================================================================

type webServiceImpl struct {
	router  *mux.Router
	ctx     context.Context
	version string
}

func (ws *webServiceImpl) PathWithVersion(prefix string) string {
	return fmt.Sprintf("/%s/%s", prefix, ws.version)
}

//create public routes
func (ws *webServiceImpl) Pub() {
	//create a subroute
	r := ws.router.PathPrefix(ws.PathWithVersion("pub")).Subrouter()
	for _, route := range routes.Get() {
		if route.Kind == routes.RoutePubKind {
			r.Handle(route.Path, route.Handler).Methods(route.Method)
		}
	}
	//
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

func renderReact(w io.Writer) error {
	f, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		return err
	}
	_, err = w.Write(f)
	if err != nil {
		return err
	}
	// success
	return nil
}

func HandlerRenderReact(w http.ResponseWriter, r *http.Request) {
	err := renderReact(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

// create privates routes
func (ws *webServiceImpl) Priv() {
	//create a subroute
	r := ws.router.PathPrefix(ws.PathWithVersion("priv")).Subrouter()
	for _, route := range routes.Get() {
		if route.Kind == routes.RoutePrivKind {
			r.Handle(route.Path, route.Handler).Methods(route.Method)
		}
	}
	//
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

func (ws *webServiceImpl) React() {
	prefix := ws.PathWithVersion("web")

	r := ws.router.PathPrefix(prefix).Subrouter()
	r.PathPrefix("/public/").Handler(
		http.StripPrefix(fmt.Sprintf("%s/public/", prefix), http.FileServer(http.Dir("./web"))),
	)
	r.PathPrefix("/react/").HandlerFunc(
		HandlerRenderReact,
	)
}

// initialize routess
func (ws *webServiceImpl) Init() {
	ws.Pub()
	ws.Priv()
	ws.React()
}

// return routes handler
func (ws *webServiceImpl) GetRouters() http.Handler {
	return (ws.router)
}

//==============================================================================
//  Static functions
//==============================================================================

// NewWebService create a new instance of WebService
func NewWebService(ctx context.Context, version string) WebService {
	return &webServiceImpl{
		router:  mux.NewRouter(),
		ctx:     ctx,
		version: version,
	}
}
