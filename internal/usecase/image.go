package usecase

import (
	"context"
	"fmt"
	"github.com/degeboman/betera-test-task/internal/gateway"
	"github.com/degeboman/betera-test-task/internal/models"
	"net/http"
)

type ImageSaveUseCase interface {
	Apply(models.ImageUnitCore) (string, error)
}

type ImageByNameUseCase interface {
	Apply(imageName string) (models.ImageUnitCore, error)
}

type ImageDownloadFromUrlUseCase interface {
	Apply(url string) (models.ImageUnitCore, error)
}

type imageSaver interface {
	Save(context.Context, models.ImageUnitCore) (string, error)
}

type ImageSaveUseCaseImpl struct {
	imageSaver
}

func NewImageSaveUseCaseImpl(imageSaver gateway.Gateway) *ImageSaveUseCaseImpl {
	return &ImageSaveUseCaseImpl{imageSaver: imageSaver}
}

func (i ImageSaveUseCaseImpl) Apply(imageCore models.ImageUnitCore) (string, error) {
	return i.imageSaver.Save(context.TODO(), imageCore)
}

type imageGetterByName interface {
	ByName(ctx context.Context, imageName string) (models.ImageUnitCore, error)
}

type ImageByNameUseCaseImpl struct {
	imageGetterByName
}

func NewImageByNameUseCaseImpl(imageGetterByName gateway.Gateway) *ImageByNameUseCaseImpl {
	return &ImageByNameUseCaseImpl{imageGetterByName: imageGetterByName}
}

type ImageDownloadFromUrlUseCaseImpl struct{}

func NewImageDownloadFromUrlUseCaseImpl() *ImageDownloadFromUrlUseCaseImpl {
	return &ImageDownloadFromUrlUseCaseImpl{}
}

func (i ImageDownloadFromUrlUseCaseImpl) Apply(url string) (models.ImageUnitCore, error) {
	const op = "usecase.image.DownloadFromUrl"

	response, err := http.Get(url)
	if err != nil {
		return models.ImageUnitCore{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.ImageUnitCore{
		Payload:     response.Body,
		PayloadSize: response.ContentLength,
	}, nil
}

func (i ImageByNameUseCaseImpl) Apply(imageName string) (models.ImageUnitCore, error) {
	return i.imageGetterByName.ByName(context.TODO(), imageName)
}
