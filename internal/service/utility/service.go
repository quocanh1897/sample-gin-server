package utilityusecase

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
	apperror "github.com/quocanh1897/sample-gin-server/internal/error"
	dto "github.com/quocanh1897/sample-gin-server/internal/model/dto/response"
	"github.com/quocanh1897/sample-gin-server/internal/pkg/logger"
	"github.com/quocanh1897/sample-gin-server/internal/repository"
	"github.com/quocanh1897/sample-gin-server/internal/utils"
)

//go:generate mockery --name=UtilityService --case=snake
type Service interface {
	HealthCheck(ctx context.Context) dto.HealthCheck
	GetTimeZones(ctx context.Context) ([]dto.SystemTimeZone, error)
}

type utilityServiceImpl struct {
	logger    *slog.Logger
	appConfig config.AppConfig
}

func NewUtilityService() Service {
	return &utilityServiceImpl{
		logger:    logger.NewLoggerWithClassName("utilityServiceImpl"),
		appConfig: config.GetAppConfig(),
	}
}

func (u *utilityServiceImpl) HealthCheck(ctx context.Context) dto.HealthCheck {
	repositories := []repository.Interface{}

	var wg sync.WaitGroup
	wg.Add(len(repositories))
	resultChan := make(chan dto.HealthCheck)

	for _, repo := range repositories {
		go u.healthCheckDependencyAsync(ctx, repo, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	res := dto.NewOkHealthCheck("sgs", nil)
	for result := range resultChan {
		res.Dependencies = append(res.Dependencies, result)
		if result.Status == constant.ServiceStatusNotAvailable {
			res.Status = constant.ServiceStatusNotAvailable
		}
	}

	return res
}

func (u *utilityServiceImpl) GetTimeZones(ctx context.Context) ([]dto.SystemTimeZone, error) {
	result, err := utils.LoadListTimeZone([]byte(constant.TimeZones))
	if err != nil {
		u.logger.ErrorContext(ctx, fmt.Sprintf("[GetTimeZones] Fail to load list of time zone, reason: %s", err))
		return []dto.SystemTimeZone{}, apperror.NewServiceUnavailableError("Fail to get list timezone")
	}
	return result, nil
}

func (u *utilityServiceImpl) healthCheckDependencyAsync(
	ctx context.Context,
	repo repository.Interface,
	wg *sync.WaitGroup,
	resultChan chan dto.HealthCheck,
) {
	defer wg.Done()
	if err := repo.HealthCheck(ctx); err != nil {
		resultChan <- dto.NewFailHealthCheck(repo.GetName(), err.Error(), nil)
	} else {
		resultChan <- dto.NewOkHealthCheck(repo.GetName(), nil)
	}
}
