package routes

import "odonto/internal/interf/resource"

func init() {
	Registry([]Route{
		{
			Kind:    RoutePrivKind,
			Path:    "/doctor",
			Method:  "POST",
			Handler: resource.CreateDoctorHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/doctors",
			Method:  "GET",
			Handler: resource.ListDoctorHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/doctor/{personPid}",
			Method:  "GET",
			Handler: resource.FindDoctorHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/doctor/{personPid}",
			Method:  "PUT",
			Handler: resource.UpdateDoctorHandler(),
		},
		{
			Kind:    RoutePrivKind,
			Path:    "/doctor/{personPid}",
			Method:  "DELETE",
			Handler: resource.RemoveDoctorHandler(),
		},
	})
}

