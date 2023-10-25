package s3_storage

import (
	"context"
	"github.com/degeboman/betera-test-task/internal/models"
)

type S3Storage interface {
	UploadFile(context.Context, models.ImageUnit) (string, error)
	DownloadFile(context.Context, string) (models.ImageUnit, error)
}
