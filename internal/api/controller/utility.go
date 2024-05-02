package controller

import (
	"log/slog"
	"net/http"

	"git.taservs.net/platform/edca-api/internal/constant"
	dto "git.taservs.net/platform/edca-api/internal/model/dto/response"
	"git.taservs.net/platform/edca-api/internal/pkg/logger"
	utilityusecase "git.taservs.net/platform/edca-api/internal/usecase/utilitiy"
	"github.com/gin-gonic/gin"
)

//go:generate mockery --name=UtilityController --case=snake
type UtilityController interface {
	HealthCheck(ctx *gin.Context)
	GetTimeZones(ctx *gin.Context)
}

func NewUtilityController(utilityUseCase utilityusecase.UtilityUseCase) UtilityController {
	return &utilityControllerImpl{
		utilityUseCase,
		logger.NewLoggerWithClassName("utilityControllerImpl"),
	}
}

type utilityControllerImpl struct {
	utilityUseCase utilityusecase.UtilityUseCase
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
