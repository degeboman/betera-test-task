package handlers

import (
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	mwLogger "github.com/degeboman/betera-test-task/internal/http-server/middleware/logger"
	"github.com/degeboman/betera-test-task/internal/logger/sl"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/degeboman/betera-test-task/internal/usecase"
	"github.com/degeboman/betera-test-task/pkg/lib/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	log         *slog.Logger
	apodUsecase usecase.ApodUsecase
}

func New(logger *slog.Logger, usecase usecase.Usecase) Handler {
	return Handler{
		logger,
		usecase.ApodUsecase,
	}
}

func (h Handler) SetupRouter(log *slog.Logger) *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route(constant.RouteApod, func(r chi.Router) {
		r.Get(constant.RouteAll, h.All())
		r.Get(constant.RouteByDate, h.ByDate())
	})

	return router
}

type Response struct {
}

func (h Handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.apod.all.All"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		_, err := h.apodUsecase.All()

		if err != nil {
			log.Error("failed to get all apod", sl.Err(err))

			render.JSON(w, r, api.Error("failed to sign up"+sl.Err(err).String()))

			return
		}

		render.JSON(w, r, Response{})
	}
}

func (h Handler) ByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.handler.ByDate"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		date := chi.URLParam(r, "date")

		if !validateApodDate(date) {
			log.Error("failed to validate date")

			render.JSON(w, r, api.Error("failed to validate date"))

			return
		}

		apod, err := h.apodUsecase.GetByDate(date)
		if err != nil {
			log.Error("failed to get apod by date", sl.Err(err))

			render.JSON(w, r, api.Error("failed to sign up"+sl.Err(err).String()))

			return
		}

		render.HTML(w, r, getHtmlResponse(apod))
	}
}

func getHtmlResponse(apod models.ApodCore) string {
	return "<html> " +
		fmt.Sprintf("<h1>%s</h1>", apod.Title) +
		fmt.Sprintf("<h3>%s</h3>", apod.Date) +
		fmt.Sprintf("<img src=\"%s\" />", apod.Url) +
		fmt.Sprintf("<p>%s</p>", apod.Explanation) +
		" </html>"
}

func validateApodDate(date string) bool {

	const layout = "2006-01-02"

	parseDate, err := time.Parse(layout, date)
	if err != nil {
		return false
	} else {
		// checking for a date from the future
		if parseDate.After(time.Now()) {
			return false
		}

		return true
	}
}
