package storage

import (
	"github.com/degeboman/betera-test-task/internal/models"
)

type ApodStorage interface {
	Create(gorm models.ApodGorm) error
	ByDate(date string) (models.ApodGorm, error)
	All() ([]models.ApodGorm, error)
}
