package main

import (
	apodworker "github.com/degeboman/betera-test-task/internal/apod-worker"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/gateway"
	httpserver "github.com/degeboman/betera-test-task/internal/http-server"
	"github.com/degeboman/betera-test-task/internal/http-server/handlers"
	"github.com/degeboman/betera-test-task/internal/logger"
	minioclient "github.com/degeboman/betera-test-task/internal/s3-storage/minio-client"
	postgresclient "github.com/degeboman/betera-test-task/internal/storage/postgres-client"
	"github.com/degeboman/betera-test-task/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	app := invokeWith(fx.Invoke(httpserver.New))
	app.Run()
}

func invokeWith(options ...fx.Option) *fx.App {

	// Provide registers any number of constructor functions, teaching the application how to instantiate various types.
	// The supplied constructor function(s) may depend on other types available in the application, must return
	// one or more objects, and may return an error.

	di := []fx.Option{
		fx.Provide(config.MustLoad),
		fx.Provide(logger.SetupLogger),
		fx.Provide(postgresclient.MustLoad),
		fx.Provide(minioclient.MustLoad),
		fx.Provide(gateway.SetupGateway),
		fx.Provide(usecase.NewApodGetByDateUseCaseImpl),
		fx.Provide(usecase.NewApodCreateUseCaseImpl),
		fx.Provide(usecase.NewApodUploadByDateUseCaseImpl),
		fx.Provide(usecase.NewApodAllUseCaseImpl),
		fx.Provide(usecase.NewImageSaveUseCaseImpl),
		fx.Provide(usecase.NewImageByNameUseCaseImpl),
		fx.Provide(usecase.NewImageDownloadFromUrlUseCaseImpl),
		fx.Provide(handlers.SetupRouter),
		fx.Provide(func() chan bool { return make(chan bool) }), // chan for stopping worker
		fx.Provide(apodworker.New),
	}

	for _, option := range options {
		di = append(di, option)
	}

	return fx.New(di...)
}
