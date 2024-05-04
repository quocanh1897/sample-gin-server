package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
	apperror "github.com/quocanh1897/sample-gin-server/internal/error"
	sessionservice "github.com/quocanh1897/sample-gin-server/internal/service/session"
	"net/http"
)

type SessionValidatorMiddleware struct {
	config         *config.AppConfig
	sessionService sessionservice.Service
}

func SessionValidatorHandler(
	cfg *config.AppConfig,
	sessionService sessionservice.Service,
) gin.HandlerFunc {
	return SessionValidatorMiddleware{
		config:         cfg,
		sessionService: sessionService,
	}.HandlerFunc
}

func (m SessionValidatorMiddleware) HandlerFunc(ctx *gin.Context) {
	authorization, err := m.sessionService.Authorize(ctx, *ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, apperror.NewUnauthorizedError("Unauthorized"))
		return
	}

	var scopes []constant.Scope
	for _, v := range authorization.Scopes {
		scopes = append(scopes, constant.Scope(v))
	}

	ctx.Set(constant.SessionScopes, scopes)
	ctx.Next()
}
