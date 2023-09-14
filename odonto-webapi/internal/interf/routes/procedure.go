package routes

import "odonto/internal/interf/resource"

func init() {
	Registry([]Route{
		{
			Kind:    RoutePrivKind,
			Path:    "/procedure",
			Method:  "POST",
			Handler: resource.CreateProcedureHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/procedures",
			Method:  "GET",
			Handler: resource.ListProcedureHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/procedure/{pid}",
			Method:  "GET",
			Handler: resource.FindProcedureHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/procedure/{pid}",
			Method:  "PUT",
			Handler: resource.UpdateProcedureHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/procedure/{pid}",
			Method:  "DELETE",
			Handler: resource.RemoveProcedureHandler(),
		},
	})
}

