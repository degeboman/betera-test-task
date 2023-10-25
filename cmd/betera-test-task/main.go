package main

import (
	apodclient "github.com/degeboman/betera-test-task/internal/apod-client"
	apodworker "github.com/degeboman/betera-test-task/internal/apod-worker"
	"github.com/degeboman/betera-test-task/internal/config"
	httpserver "github.com/degeboman/betera-test-task/internal/http-server"
	"github.com/degeboman/betera-test-task/internal/http-server/handlers"
	"github.com/degeboman/betera-test-task/internal/logger"
	minioclient "github.com/degeboman/betera-test-task/internal/s3-storage/minio-client"
	"github.com/degeboman/betera-test-task/internal/storage"
	postgresclient "github.com/degeboman/betera-test-task/internal/storage/postgres-client"
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
		fx.Provide(apodclient.New),
		fx.Provide(storage.New),
		fx.Provide(func() chan bool { return make(chan bool) }),
		fx.Provide(apodworker.New),
		fx.Provide(handlers.New),
	}

	for _, option := range options {
		di = append(di, option)
	}

	return fx.New(di...)
}
