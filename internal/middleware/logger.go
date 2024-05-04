package middleware

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
)

func Logger() gin.HandlerFunc {
	return LoggerMiddleware{}.HandlerFunc
}

type LoggerMiddleware struct{}

func (m LoggerMiddleware) HandlerFunc(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()

	param := m.prepareParam(ctx, start)
	msg := m.prepareMessage(param)

	if param.StatusCode < 400 {
		m.injectFields(param).InfoContext(ctx, msg)
	} else if param.StatusCode < 500 {
		m.injectFields(param).WarnContext(ctx, msg)
	} else {
		m.injectFields(param).ErrorContext(ctx, msg)
	}
}

func (m LoggerMiddleware) prepareParam(ctx *gin.Context, start time.Time) gin.LogFormatterParams {
	param := gin.LogFormatterParams{
		Request: ctx.Request,
		Keys:    ctx.Keys,
	}

	param.TimeStamp = time.Now()
	param.Latency = param.TimeStamp.Sub(start)

	param.ClientIP = ctx.ClientIP()
	param.Method = ctx.Request.Method
	param.StatusCode = ctx.Writer.Status()
	param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()

	param.BodySize = ctx.Writer.Size()

	path := ctx.Request.URL.Path
	if ctx.Request.URL.RawQuery != "" {
		path += "?" + ctx.Request.URL.RawQuery
	}

	param.Path = path

	return param
}

func (m LoggerMiddleware) prepareMessage(param gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"%s %s %s %d %dms %s",
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Milliseconds(),
		param.Request.UserAgent(),
	)
}

func (m LoggerMiddleware) injectFields(param gin.LogFormatterParams) *slog.Logger {
	return slog.With(
		slog.Int(constant.StatusCodeAttributeKey, param.StatusCode),
		slog.Int64(constant.LatencyAttributeKey, param.Latency.Milliseconds()),
		slog.String(constant.ClientIPAttributeKey, param.ClientIP),
		slog.String(constant.MethodAttributeKey, param.Method),
		slog.String(constant.PathAttributeKey, param.Path),
		slog.String(constant.ErrorMessageAttributeKey, param.ErrorMessage),
	)
}
