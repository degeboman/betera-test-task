package storage

import (
	"errors"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/models"
	postgresclient "github.com/degeboman/betera-test-task/internal/storage/postgres-client"
	"gorm.io/gorm"
)

type ApodStorage interface {
	Create(gorm models.ApodGorm) error
	ByDate(date string) (models.ApodGorm, error)
	All() ([]models.ApodGorm, error)
}

type ApodStorageImpl struct {
	Pc postgresclient.PostgresClient
}

func New(pc postgresclient.PostgresClient) ApodStorageImpl {
	return ApodStorageImpl{
		Pc: pc,
	}
}

func (a ApodStorageImpl) Create(apod models.ApodGorm) error {
	const op = "storage.storage.Create"

	if err := a.Pc.Db.Create(&apod).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a ApodStorageImpl) ByDate(date string) (models.ApodGorm, error) {
	const op = "storage.storage.ByDate"

	var apodGorm models.ApodGorm

	if err := a.Pc.Db.Where("date = ?", date).Take(&apodGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apodGorm, errors.New(constant.ErrRecordNotFound)
		}

		return apodGorm, fmt.Errorf("%s: %w", op, err)
	}

	return apodGorm, nil
}

func (a ApodStorageImpl) All() ([]models.ApodGorm, error) {
	const op = "storage.storage.All"
	var apodsGorm []models.ApodGorm

	if err := a.Pc.Db.Find(&apodsGorm).Error; err != nil {
		return apodsGorm, fmt.Errorf("%s: %w", op, err)
	}

	return apodsGorm, nil
}
