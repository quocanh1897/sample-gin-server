package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
)

type customHandler struct {
	handler slog.Handler
}

func (h *customHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &customHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return &customHandler{handler: h.handler.WithGroup(name)}
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	if c, ok := ctx.(*gin.Context); ok {
		r.AddAttrs(
			slog.String(constant.RequestIdAttributeKey, c.GetHeader(constant.RequestIdHeaderKey)),
			slog.String(constant.RequestedWithAttributeKey, c.GetHeader(constant.RequestedWithHeaderKey)),
		)
	}
	return h.handler.Handle(ctx, r)
}

func Init() {
	cfg := config.GetLoggingConfig()

	var logHandler slog.Handler
	logLevel := new(slog.LevelVar)
	handlerOption := slog.HandlerOptions{Level: logLevel}

	switch cfg.Formatter {
	case config.LoggingJSONFormatter:
		logHandler = &customHandler{
			handler: slog.NewJSONHandler(os.Stdout, &handlerOption),
		}
	case config.LoggingTextFormatter:
		logHandler = &customHandler{
			handler: slog.NewTextHandler(os.Stdout, &handlerOption),
		}
	default:
		logHandler = &customHandler{
			handler: slog.NewJSONHandler(os.Stdout, &handlerOption),
		}
	}

	switch cfg.Level {
	case config.LoggingInfoLevel:
		logLevel.Set(slog.LevelInfo)
	case config.LoggingDebugLevel:
		logLevel.Set(slog.LevelDebug)
	default:
		logLevel.Set(slog.LevelInfo)
	}

	logger := slog.New(logHandler).With(constant.ServiceNameAttributeKey, constant.ServiceNameValue)
	slog.SetDefault(logger)
}

func NewLoggerWithClassName(className string) *slog.Logger {
	return slog.With(constant.ClassNameAttributeKey, className)
}
