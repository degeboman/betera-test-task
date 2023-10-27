package handlers

import (
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/http-server/handlers/all"
	bydate "github.com/degeboman/betera-test-task/internal/http-server/handlers/by-date"
	mwlogger "github.com/degeboman/betera-test-task/internal/http-server/middleware/logger"
	"github.com/degeboman/betera-test-task/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

func SetupRouter(log *slog.Logger,
	apodGetByDateUseCase *usecase.ApodGetByDateUseCaseImpl,
	apodAllUseCase *usecase.ApodAllUseCaseImpl,
) *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwlogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route(constant.RouteApod, func(r chi.Router) {
		r.Get(constant.RouteAll, all.New(log, apodAllUseCase))
		r.Get(constant.RouteByDate, bydate.New(log, apodGetByDateUseCase))
	})

	return router
}
