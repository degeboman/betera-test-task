package usecase

import (
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/gateway"
)

// Usecase is business logic layer
type Usecase struct {
	ApodUsecase
	ImageUsecase
}

func SetupUsecase(cfg config.Config, gateway gateway.Gateway) Usecase {
	return Usecase{
		ApodUsecase: ApodUsecaseImpl{
			cfg,
			gateway,
		},
		ImageUsecase: ImageUsecaseImpl{
			gateway.S3Storage,
		},
	}
}
