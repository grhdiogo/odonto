package routes

func init() {
	Registry([]Route{
		// {
		// 	Kind:    RoutePubKind,
		// 	Path:    "/auth",
		// 	Method:  http.MethodPost,
		// 	Handler: resource.CreateAuthaccessHandler(),
		// },
		// {
		// 	Kind:    RoutePrivKind,
		// 	Path:    "/check-token",
		// 	Method:  http.MethodGet,
		// 	Handler: resource.CheckTokenHandler(),
		// },
		// {
		// 	Kind:    RoutePrivKind,
		// 	Path:    "/logout",
		// 	Method:  http.MethodDelete,
		// 	Handler: resource.RemoveAuthaccessHandler(),
		// },
		// {
		// 	Kind:    RoutePrivKind,
		// 	Path:    "/verify-first-access",
		// 	Method:  http.MethodPost,
		// 	Handler: resource.VerifyFirstAccessHandler(),
		// },
	})
}
