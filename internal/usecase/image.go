package usecase

import (
	"context"
	"fmt"
	"github.com/degeboman/betera-test-task/internal/models"
	s3storage "github.com/degeboman/betera-test-task/internal/s3-storage"
	"net/http"
)

type ImageUsecase interface {
	SaveImage(models.ImageUnit) (string, error)
	ByName(imageName string) (models.ImageUnit, error)
	DownloadFromUrl(url string) (models.ImageUnit, error)
}

type ImageUsecaseImpl struct {
	s3storage.S3Storage
}

func (i ImageUsecaseImpl) DownloadFromUrl(url string) (models.ImageUnit, error) {
	const op = "usecase.image.DownloadFromUrl"

	response, err := http.Get(url)
	if err != nil {
		return models.ImageUnit{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.ImageUnit{
		Payload:     response.Body,
		PayloadSize: response.ContentLength,
	}, nil
}

func (i ImageUsecaseImpl) SaveImage(unit models.ImageUnit) (string, error) {
	return i.S3Storage.UploadImage(context.TODO(), unit)
}

func (i ImageUsecaseImpl) ByName(imageName string) (models.ImageUnit, error) {
	return i.S3Storage.DownloadImage(context.TODO(), imageName)
}
