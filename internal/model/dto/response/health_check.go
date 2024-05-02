package dto

import "github.com/quocanh1897/sample-gin-server/internal/constant"

type HealthCheck struct {
	Service      string                 `json:"service"`
	Status       constant.ServiceStatus `json:"status"`
	Message      string                 `json:"message"`
	Dependencies []HealthCheck          `json:"dependencies"`
}

func NewFailHealthCheck(service string, msg string, dependencies []HealthCheck) HealthCheck {
	return HealthCheck{
		Service:      service,
		Status:       constant.ServiceStatusNotAvailable,
		Message:      msg,
		Dependencies: dependencies,
	}
}

func NewOkHealthCheck(service string, dependencies []HealthCheck) HealthCheck {
	return HealthCheck{
		Service:      service,
		Status:       constant.ServiceStatusAvailable,
		Dependencies: dependencies,
	}
}
