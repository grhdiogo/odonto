package routes

import "net/http"

type RouteKind int

const (
	RoutePubKind RouteKind = 0
	RoutePrivKind RouteKind = 1
)

type Route struct {
	Kind    RouteKind
	Path    string
	Method  string
	Handler http.Handler
}

var defaultRoutes = make([]Route, 0)

func Registry(rs []Route) {
	defaultRoutes = append(defaultRoutes, rs...)
}

func Get() []Route {
	return defaultRoutes
}
