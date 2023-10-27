package s3_storage

import (
	"context"
	"github.com/degeboman/betera-test-task/internal/models"
)

type S3Storage interface {
	Save(context.Context, models.ImageUnitCore) (string, error)
	ByName(context.Context, string) (models.ImageUnitCore, error)
}
