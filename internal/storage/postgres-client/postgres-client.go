package postgres_client

import (
	"errors"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type PostgresClient struct {
	Db *gorm.DB
}

func MustLoad(cfg config.Config) PostgresClient {
	db, err := gorm.Open(
		postgres.Open(cfg.PostgresDsn),
		&gorm.Config{Logger: gormLogger.New(
			log.New(os.Stdout, "[GORM]\t", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  gormLogger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		)},
	)

	if err != nil {
		log.Fatalf("failed to open db session: %s", err.Error())
	}

	pc := PostgresClient{
		Db: db,
	}

	if err := pc.Migrate(); err != nil {
		log.Fatalf("failed to migrate: %s", err.Error())
	}

	return pc
}

func (p PostgresClient) Migrate() error {
	return p.Db.AutoMigrate(
		&models.ApodGorm{},
	)
}

func (p PostgresClient) Create(apod models.ApodGorm) error {
	const op = "storage.storage.Create"

	if err := p.Db.Create(&apod).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (p PostgresClient) ByDate(date string) (models.ApodGorm, error) {
	const op = "storage.storage.ByDate"

	var apodGorm models.ApodGorm

	if err := p.Db.Where("date = ?", date).Take(&apodGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apodGorm, errors.New(constant.ErrRecordNotFound)
		}

		return apodGorm, fmt.Errorf("%s: %w", op, err)
	}

	return apodGorm, nil
}

func (p PostgresClient) All() ([]models.ApodGorm, error) {
	const op = "storage.storage.All"
	var apodsGorm []models.ApodGorm

	if err := p.Db.Find(&apodsGorm).Error; err != nil {
		return apodsGorm, fmt.Errorf("%s: %w", op, err)
	}

	return apodsGorm, nil
}
