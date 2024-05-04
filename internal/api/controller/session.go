package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"

	"github.com/quocanh1897/sample-gin-server/internal/pkg/logger"

	sessionservice "github.com/quocanh1897/sample-gin-server/internal/service/session"
)

type SessionController interface {
	GetSession(ctx *gin.Context)
}

func NewSessionController(sessionService sessionservice.Service) SessionController {
	return &SessionControllerImpl{
		logger.NewLoggerWithClassName("SessionControllerImpl"),
		sessionService,
	}
}

type SessionControllerImpl struct {
	logger         *slog.Logger
	sessionService sessionservice.Service
}

// GetSession implements SessionController.
func (session *SessionControllerImpl) GetSession(ctx *gin.Context) {
	logger.NewLoggerWithClassName("SessionControllerImpl").Debug("GetSession")
	ctx.AddParam("session-abc", "session-abc")
}
