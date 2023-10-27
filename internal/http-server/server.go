package http_server

import (
	"context"
	apodworker "github.com/degeboman/betera-test-task/internal/apod-worker"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/logger/sl"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(
	stopWorker chan bool,
	lifecycle fx.Lifecycle,
	log *slog.Logger,
	cfg config.Config,
	worker apodworker.ApodWorker,
	router *chi.Mux,
) {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.HttpServerAddress,
		Handler:      router,
		ReadTimeout:  cfg.HttpServerTimeout,
		WriteTimeout: cfg.HttpServerTimeout,
		IdleTimeout:  cfg.HttpServerIdleTimeout,
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				log.Info("starting server")

				go worker.Ticker(stopWorker)

				go func() {
					if err := srv.ListenAndServe(); err != nil {
						log.Error("failed to start server", sl.Err(err))
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("stopping server")

				<-done
				stopWorker <- true

				timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				if err := srv.Shutdown(timeout); err != nil {
					log.Error("failed to stop server", sl.Err(err))

					return err
				}

				log.Info("server stopped")

				return nil
			},
		},
	)
}
