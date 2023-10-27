package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/gateway"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/degeboman/betera-test-task/internal/models/mapping"
	"github.com/degeboman/betera-test-task/internal/storage"
	"io"
	"net/http"
)

type apodGetterByDate interface {
	ByDate(date string) (models.ApodGorm, error)
}

type apodCreator interface {
	Create(gorm models.ApodGorm) error
}

type apodGetterAll interface {
	All() ([]models.ApodGorm, error)
}

type ApodCreateUseCase interface {
	Apply(apod models.ApodCore) error
}

type ApodGetByDateUseCase interface {
	Apply(date string) (models.ApodCore, error)
}

type ApodUploadByDateUseCase interface {
	Apply(date string) (models.ApodCore, error)
}

type ApodAllUseCase interface {
	Apply() ([]models.ApodCore, error)
}

type ApodGetByDateUseCaseImpl struct {
	apodGetterByDate
}

func NewApodGetByDateUseCaseImpl(apodGetterByDate gateway.Gateway) *ApodGetByDateUseCaseImpl {
	return &ApodGetByDateUseCaseImpl{apodGetterByDate: apodGetterByDate}
}

type ApodCreateUseCaseImpl struct {
	apodCreator
}

func NewApodCreateUseCaseImpl(apodCreator gateway.Gateway) *ApodCreateUseCaseImpl {
	return &ApodCreateUseCaseImpl{apodCreator: apodCreator}
}

type ApodUploadByDateUseCaseImpl struct {
	cfg config.Config
}

func NewApodUploadByDateUseCaseImpl(cfg config.Config) *ApodUploadByDateUseCaseImpl {
	return &ApodUploadByDateUseCaseImpl{cfg: cfg}
}

type ApodAllUseCaseImpl struct {
	apodGetterAll
}

func NewApodAllUseCaseImpl(apodGetterAll gateway.Gateway) *ApodAllUseCaseImpl {
	return &ApodAllUseCaseImpl{apodGetterAll: apodGetterAll}
}

func (a ApodAllUseCaseImpl) Apply() ([]models.ApodCore, error) {
	const op = "usecase.apod.All"

	apodsGorm, err := a.apodGetterAll.All()
	if err != nil {
		return []models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}

	apodsCore := make([]models.ApodCore, 0, len(apodsGorm))

	for _, apodGorm := range apodsGorm {
		apodsCore = append(apodsCore, mapping.ApodGormToCore(apodGorm))
	}

	return apodsCore, nil
}

func (a ApodUploadByDateUseCaseImpl) Apply(date string) (models.ApodCore, error) {
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

	var aw models.ApodWeb

	if err := json.Unmarshal(body, &aw); err != nil {
		return models.ApodCore{}, fmt.Errorf("%s: %w", op, err)
	}

	apodCore := mapping.ApodWebToCore(aw)

	return apodCore, nil
}

func (a ApodGetByDateUseCaseImpl) Apply(date string) (models.ApodCore, error) {

	apodGorm, err := a.apodGetterByDate.ByDate(date)
	if err != nil {
		return models.ApodCore{}, err
	}

	apodCore := mapping.ApodGormToCore(apodGorm)

	return apodCore, nil
}

func (a ApodCreateUseCaseImpl) Apply(apod models.ApodCore) error {

	apodGorm := mapping.ApodCoreToGorm(apod)

	return a.apodCreator.Create(apodGorm)
}

type ApodUsecaseImpl struct {
	cfg config.Config
	storage.ApodStorage
}

func apodByDateUrl(apiKey, date string) string {
	return constant.Url + "?api_key=" + apiKey + "&date=" + date
}
