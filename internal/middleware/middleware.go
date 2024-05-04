package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/quocanh1897/sample-gin-server/internal/service/session"

	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
)

type middlewareImpl struct {
	appConfig               *config.AppConfig
	loggerHandler           gin.HandlerFunc
	errorHandlerHandler     gin.HandlerFunc
	contextHandler          gin.HandlerFunc
	sessionValidatorHandler gin.HandlerFunc
	scopeValidatorFactory   func([][]constant.Scope) gin.HandlerFunc
}

func New(appConfig *config.AppConfig, sessionService session.Service) Middleware {
	return &middlewareImpl{
		loggerHandler:           Logger(),
		errorHandlerHandler:     ErrorHandler(),
		contextHandler:          ContextHandler(appConfig),
		sessionValidatorHandler: SessionValidatorHandler(appConfig, sessionService),
		scopeValidatorFactory:   ScopeValidatorMiddlewareFactory(),
	}
}

type Factory interface {
	ScopeValidatorFactory() func([][]constant.Scope) gin.HandlerFunc
}

type Middleware interface {
	Factory
	LoggerHandler() gin.HandlerFunc
	ErrorHandlerHandler() gin.HandlerFunc
	ContextHandler() gin.HandlerFunc
	SessionValidatorHandler() gin.HandlerFunc
}

func (i *middlewareImpl) LoggerHandler() gin.HandlerFunc {
	return i.loggerHandler
}

func (i *middlewareImpl) ErrorHandlerHandler() gin.HandlerFunc {
	return i.errorHandlerHandler
}

func (i *middlewareImpl) ContextHandler() gin.HandlerFunc {
	return i.contextHandler
}

func (i *middlewareImpl) SessionValidatorHandler() gin.HandlerFunc {
	return i.sessionValidatorHandler
}

func (i *middlewareImpl) ScopeValidatorFactory() func([][]constant.Scope) gin.HandlerFunc {
	return i.scopeValidatorFactory
}
