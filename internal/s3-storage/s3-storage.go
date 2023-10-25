package s3_storage

import (
	"context"
	"github.com/degeboman/betera-test-task/internal/models"
)

type S3Storage interface {
	UploadImage(context.Context, models.ImageUnit) (string, error)
	DownloadImage(context.Context, string) (models.ImageUnit, error)
}
