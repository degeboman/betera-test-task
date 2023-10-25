package postgres_client

import (
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
