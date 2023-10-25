package config

import (
	"flag"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	PostgresDsn string
	NasaApiKey  string `yaml:"nasa_api_key" env-required:"true"`
	Env         string `yaml:"env" env-default:"local"`

	HTTPServer    `yaml:"http_server"`
	MinioAuthData `yaml:"minio_auth_data"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:1010"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type MinioAuthData struct {
	Url      string `yaml:"url"`
	User     string `yaml:"user" env-default:"minio"`
	Password string `yaml:"password" env-required:"true"`
	Ssl      bool   `yaml:"ssl" env-default:"false"`
}

//type Database struct {
//	Port     string `env:"POSTGRES_PORT" env-default:"4444"`
//	Host     string `env:"HOST" env-default:"localhost"`
//	Name     string `env:"POSTGRES_DBNAME" env-default:"apod_db"`
//	User     string `env:"POSTGRES_USER" env-default:"apod"`
//	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
//}

func MustLoad() Config {
	configPath := flag.String(
		constant.ConfigPathFlag,
		"../../config/local.yml",
		constant.ConfigPathFlagUsage,
	)

	// check if file exists
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", *configPath)
	}

	postgresDsnString := flag.String(
		constant.PostgrestDsnFlag,
		"",
		constant.PostgresDsnFlaUsage,
	)

	flag.Parse()

	if *postgresDsnString == "" {
		log.Fatal("postgres dsn string is not specified")
	}

	var cfg Config

	cfg.PostgresDsn = *postgresDsnString

	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}
