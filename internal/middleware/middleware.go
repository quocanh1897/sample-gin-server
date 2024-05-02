package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
)

type middlewareImpl struct {
	appConfig             *config.AppConfig
	loggerHandler         gin.HandlerFunc
	errorHandlerHandler   gin.HandlerFunc
	contextHandler        gin.HandlerFunc
	scopeValidatorFactory func([][]constant.Scope) gin.HandlerFunc
}

type MiddleWareFactory interface {
	ScopeValidatorFactory() func([][]constant.Scope) gin.HandlerFunc
}

type Middleware interface {
	MiddleWareFactory
	LoggerHandler() gin.HandlerFunc
	ErrorHandlerHandler() gin.HandlerFunc
	ContextHandler() gin.HandlerFunc
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

func (i *middlewareImpl) ScopeValidatorFactory() func([][]constant.Scope) gin.HandlerFunc {
	return i.scopeValidatorFactory
}

func New(appConfig *config.AppConfig) Middleware {
	return &middlewareImpl{
		loggerHandler:         Logger(),
		errorHandlerHandler:   ErrorHandler(),
		contextHandler:        ContextHandler(appConfig),
		scopeValidatorFactory: ScopeValidatorMiddlewareFactory(),
	}
}
