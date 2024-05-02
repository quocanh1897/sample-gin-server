package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
	dto "github.com/quocanh1897/sample-gin-server/internal/model/dto/response"
	"github.com/quocanh1897/sample-gin-server/internal/pkg/logger"
	service "github.com/quocanh1897/sample-gin-server/internal/service/utility"
)

//go:generate mockery --name=UtilityController --case=snake
type UtilityController interface {
	HealthCheck(ctx *gin.Context)
	GetTimeZones(ctx *gin.Context)
}

func NewUtilityController(utilityUseCase service.Service) UtilityController {
	return &utilityControllerImpl{
		utilityUseCase,
		logger.NewLoggerWithClassName("utilityControllerImpl"),
	}
}

type utilityControllerImpl struct {
	utilityUseCase service.Service
	logger         *slog.Logger
}

func (c *utilityControllerImpl) HealthCheck(ctx *gin.Context) {
	c.logger.DebugContext(ctx, "[HealthCheck] Start func")
	defer c.logger.DebugContext(ctx, "[HealthCheck] End func")

	result := c.utilityUseCase.HealthCheck(ctx)

	status := http.StatusOK
	if result.Status == constant.ServiceStatusNotAvailable {
		status = http.StatusServiceUnavailable
	}

	ctx.JSON(status, result)
}

func (c *utilityControllerImpl) GetTimeZones(ctx *gin.Context) {
	timezones, err := c.utilityUseCase.GetTimeZones(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.SuccessResponse(timezones))
}
