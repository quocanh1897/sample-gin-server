package repository

import (
	"context"
)

type Interface interface {
	HealthCheck(ctx context.Context) error
	GetName() string
}
