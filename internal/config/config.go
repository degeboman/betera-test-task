package config

import (
	"flag"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	PostgresDsn string `env:"POSTGRES_DSN" env-required:"true"`
	NasaApiKey  string `env:"NASA_API_KEY" env-required:"true"`
	Env         string `env:"ENV" env-default:"local"`

	HttpServerAddress     string        `env:"HTTP_SERVER_ADDRESS" env-default:"localhost:1010"`
	HttpServerTimeout     time.Duration `env:"HTTP_SERVER_TIMEOUT" env-default:"4s"`
	HttpServerIdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"60s"`

	MinioUrl      string `env:"MINIO_URL"`
	MinioUser     string `env:"MINIO_ROOT_USER" env-default:"minio"`
	MinioPassword string `env:"MINIO_ROOT_PASSWORD" env-required:"true"`
	MinioSsl      bool   `env:"ssl" env-default:"false"`
}

func MustLoad() Config {
	envPath := flag.String(
		constant.EnvPathFlag,
		".env",
		constant.EnvPathFlagUsage,
	)

	flag.Parse()

	// check if file exists
	if _, err := os.Stat(*envPath); os.IsNotExist(err) {
		log.Fatalf("env file does not exist: %s", *envPath)
	}

	// loading env variables
	if err := godotenv.Load(*envPath); err != nil {
		log.Fatalf("failed to load env file: %s", *envPath)
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Println(err)
	}

	return cfg
}
