package by_date_html

import (
	"fmt"
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

func New(log *slog.Logger, useCase usecase.ApodGetByDateUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.by-date.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		date := chi.URLParam(r, "date")

		if !validateApodDate(date) {
			log.Error("failed to validate date")

			render.JSON(w, r, api.Error("failed to validate date"))

			return
		}

		apod, err := useCase.Apply(date)
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
		fmt.Sprintf("<img src=\"%s\" alt=\"%s\"/>", apod.Url, apod.Title) +
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
