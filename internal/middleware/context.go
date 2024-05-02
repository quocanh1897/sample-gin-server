package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
)

type ContextMiddleware struct{}

func ContextHandler(appConfig *config.AppConfig) gin.HandlerFunc {
	return ContextMiddleware{}.HandlerFunc
}

func (m ContextMiddleware) HandlerFunc(ctx *gin.Context) {
	ctx.Set(constant.XForwardedHost, ctx.Request.Header.Get("X-Forwarded-Host"))
	ctx.Set(constant.Host, ctx.Request.Host)
	ctx.Set(constant.IpAddress, ctx.ClientIP())

	requestId := ctx.GetHeader(constant.RequestIdHeaderKey)
	if requestId == "" {
		requestId = uuid.New().String()
	}
	ctx.Set(constant.RequestIdAttributeKey, requestId)

	ctx.Set(constant.RequestedWithAttributeKey, ctx.GetHeader(constant.RequestedWithHeaderKey))
	ctx.Next()
}
