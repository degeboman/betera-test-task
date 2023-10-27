package all

import (
	"github.com/degeboman/betera-test-task/internal/logger/sl"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/degeboman/betera-test-task/internal/models/mapping"
	"github.com/degeboman/betera-test-task/internal/usecase"
	"github.com/degeboman/betera-test-task/pkg/lib/api"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Response struct {
	Apods []models.ApodWeb `json:"apods"`
}

func New(log *slog.Logger, useCase usecase.ApodAllUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.apod.all.All"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		apodsCore, err := useCase.Apply()

		if err != nil {
			log.Error("failed to get all apod", sl.Err(err))

			render.JSON(w, r, api.Error("failed to sign up"+sl.Err(err).String()))

			return
		}

		response := Response{
			Apods: make([]models.ApodWeb, 0, len(apodsCore)),
		}

		for _, apodCore := range apodsCore {
			web := mapping.ApodCoreToWeb(apodCore)
			response.Apods = append(response.Apods, web)
		}

		render.JSON(w, r, response)
	}
}
