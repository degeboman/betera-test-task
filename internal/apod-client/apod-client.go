package apod_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/degeboman/betera-test-task/constant/apod"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/models"
	s3storage "github.com/degeboman/betera-test-task/internal/s3-storage"
	minioclient "github.com/degeboman/betera-test-task/internal/s3-storage/minio-client"
	"io"
	"net/http"
)

type ApodClient struct {
	cfg       config.Config
	s3storage s3storage.S3Storage
}

func New(cfg config.Config, mc *minioclient.MinioClient) *ApodClient {
	return &ApodClient{
		cfg:       cfg,
		s3storage: mc,
	}
}

func (ac ApodClient) DownloadImage(url string) (string, error) {
	const op = "apod-client.apod-client.DownloadImage"

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer response.Body.Close()

	imageObject := models.ImageUnit{
		Payload:     response.Body,
		PayloadSize: response.ContentLength,
	}

	// save object in s3 storage
	imageName, err := ac.s3storage.UploadFile(context.TODO(), imageObject)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return imageName, nil
}

func (ac ApodClient) ByDate(date string) (*models.ApodResponse, error) {
	const op = "apod-client.apod-client.ByDate"

	response, err := http.Get(apodByDateUrl(ac.cfg.NasaApiKey, date))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var r models.ApodResponse

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &r, nil
}

func apodByDateUrl(apiKey, date string) string {
	return apod.Url + "?api_key=" + apiKey + "&date=" + date
}
