package routes

import "odonto/internal/interf/resource"

func init() {
	Registry([]Route{
		{
			Kind:    RoutePrivKind,
			Path:    "/patient",
			Method:  "POST",
			Handler: resource.CreatePatientHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/patients",
			Method:  "GET",
			Handler: resource.ListPatientHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/patient/{personPid}",
			Method:  "GET",
			Handler: resource.FindPatientHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/patient/{personPid}",
			Method:  "PUT",
			Handler: resource.UpdatePatientHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/patient/{personPid}",
			Method:  "DELETE",
			Handler: resource.RemovePatientHandler(),
		},
	})
}

