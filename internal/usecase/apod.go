package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/degeboman/betera-test-task/internal/storage"
	"io"
	"net/http"
)

type ApodUsecase interface {
	All() ([]models.ApodCore, error)
	UploadByDate(date string) (models.ApodCore, error)
	GetByDate(date string) (models.ApodCore, error)
	Create(apod models.ApodCore) error
}

type ApodUsecaseImpl struct {
	cfg config.Config
	storage.ApodStorage
}

func (a ApodUsecaseImpl) UploadByDate(date string) (models.ApodCore, error) {
	const op = "usecase.apod.UploadByDate"

	// upload model from apod api
	response, err := http.Get(apodByDateUrl(a.cfg.NasaApiKey, date))
	if err != nil {
		return models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}
	defer response.Body.Close()

	// reading response body and unmarshal
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}

	var ar models.ApodCore

	if err := json.Unmarshal(body, &ar); err != nil {
		return models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}

	return ar, nil
}

func (a ApodUsecaseImpl) Create(apod models.ApodCore) error {
	var apodGorm models.ApodGorm
	apodGorm.FromCore(&apod)

	return a.ApodStorage.Create(apodGorm)
}

func (a ApodUsecaseImpl) All() ([]models.ApodCore, error) {
	const op = "usecase.apod.All"

	apods, err := a.ApodStorage.All()
	if err != nil {
		return []models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}

	apodsCore := make([]models.ApodCore, len(apods))

	for _, apod := range apods {
		apodTemp := apod.ToCore()
		apodsCore = append(apodsCore, apodTemp)
	}

	return apodsCore, nil
}

func (a ApodUsecaseImpl) GetByDate(date string) (models.ApodCore, error) {
	apod, err := a.ApodStorage.ByDate(date)
	if err != nil {
		return models.ApodCore{}, err
	}

	return apod.ToCore(), nil
}

func apodByDateUrl(apiKey, date string) string {
	return constant.Url + "?api_key=" + apiKey + "&date=" + date
}
