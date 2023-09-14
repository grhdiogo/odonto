package routes

import "odonto/internal/interf/resource"

func init() {
	Registry([]Route{
		{
			Kind:    RoutePrivKind,
			Path:    "/appointment",
			Method:  "POST",
			Handler: resource.CreateAppointmentHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/appointments",
			Method:  "GET",
			Handler: resource.ListAppointmentHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/appointment/{aid}",
			Method:  "GET",
			Handler: resource.FindAppointmentHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/appointment/{aid}",
			Method:  "PUT",
			Handler: resource.UpdateAppointmentHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/appointment/{aid}",
			Method:  "DELETE",
			Handler: resource.RemoveAppointmentHandler(),
		},
	})
}

