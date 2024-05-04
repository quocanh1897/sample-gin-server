package session

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"

	"github.com/quocanh1897/sample-gin-server/internal/model"
	"github.com/quocanh1897/sample-gin-server/internal/pkg/logger"
)

type (
	Service interface {
		Authorize(ctx context.Context, httpCtx http.Request) (model.Authorization, error)
	}

	sessionServiceImpl struct {
		logger *slog.Logger
	}
)

func (s sessionServiceImpl) Authorize(ctx context.Context, httpCtx http.Request) (model.Authorization, error) {
	logger.NewLoggerWithClassName("sessionServiceImpl").Info("Authorize")
	mockUUID := uuid.New()
	return model.Authorization{
		Jwt:                "",
		Scopes:             []string{"org:read", "org:write", "profile:read", "profile:write"},
		OauthApplicationId: &mockUUID,
	}, nil
}

func NewSessionService() Service {
	return &sessionServiceImpl{
		logger: logger.NewLoggerWithClassName("sessionServiceImpl"),
	}
}
