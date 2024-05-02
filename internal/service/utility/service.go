package utilityusecase

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/quocanh189/internal/config"
	"github.com/quocanh189/internal/constant"
	edcaerror "github.com/quocanh189/internal/error"
	dto "github.com/quocanh189/internal/model/dto/response"
	"github.com/quocanh189/internal/pkg/logger"
	"github.com/quocanh189/internal/repository"
	"github.com/quocanh189/internal/utils"
)

//go:generate mockery --name=UtilityUseCase --case=snake
type UtilityUseCase interface {
	HealthCheck(ctx context.Context) dto.HealthCheck
	GetTimeZones(ctx context.Context) ([]dto.SystemTimeZone, error)
}

type utilityUseCaseImpl struct {
	logger    *slog.Logger
	appConfig config.AppConfig
}

func NewUtilityUseCase() UtilityUseCase {
	return &utilityUseCaseImpl{
		logger:    logger.NewLoggerWithClassName("utilityUseCaseImpl"),
		appConfig: config.GetAppConfig(),
	}
}

func (u *utilityUseCaseImpl) HealthCheck(ctx context.Context) dto.HealthCheck {
	repositories := []repository.RepositoryInterface{}

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

	res := dto.NewOkHealthCheck("edca", nil)
	for result := range resultChan {
		res.Dependencies = append(res.Dependencies, result)
		if result.Status == constant.ServiceStatusNotAvailable {
			res.Status = constant.ServiceStatusNotAvailable
		}
	}

	return res
}

func (u *utilityUseCaseImpl) GetTimeZones(ctx context.Context) ([]dto.SystemTimeZone, error) {
	result, err := utils.LoadListTimeZone([]byte(constant.TimeZones))
	if err != nil {
		u.logger.ErrorContext(ctx, fmt.Sprintf("[GetTimeZones] Fail to load list of time zone, reason: %s", err))
		return []dto.SystemTimeZone{}, edcaerror.NewServiceUnavailableError("Fail to get list timezone")
	}
	return result, nil
}

func (u *utilityUseCaseImpl) healthCheckDependencyAsync(
	ctx context.Context,
	repo repository.RepositoryInterface,
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
