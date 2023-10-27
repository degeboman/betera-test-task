package minio_client

import (
	"context"
	"fmt"
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/config"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

type MinioClient struct {
	minioAuthData
	*minio.Client
}

type minioAuthData struct {
	url      string
	user     string
	password string
	token    string
	ssl      bool
}

func MustLoad(cfg config.Config) *MinioClient {
	mc := &MinioClient{
		minioAuthData: minioAuthData{
			password: cfg.MinioAuthData.Password,
			url:      cfg.MinioAuthData.Url,
			user:     cfg.MinioAuthData.User,
			ssl:      cfg.MinioAuthData.Ssl,
		},
	}

	var err error

	mc.Client, err = minio.New(
		mc.url,
		&minio.Options{
			Creds:  credentials.NewStaticV4(mc.user, mc.password, ""),
			Secure: mc.ssl,
		},
	)

	if err != nil {
		log.Fatalf("failed to connect minio client: %s", err.Error())
	}

	return mc
}

func (m *MinioClient) Save(ctx context.Context, object models.ImageUnitCore) (string, error) {
	const op = "s3-storage.minio-client.minio-client.UploadFile"

	imageName := generateObjectName()

	_, err := m.Client.PutObject(
		ctx,
		constant.BucketName,
		imageName,
		object.Payload,
		object.PayloadSize,
		minio.PutObjectOptions{ContentType: "image/jpeg"},
	)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return imageName, err
}

func (m *MinioClient) ByName(ctx context.Context, objectName string) (models.ImageUnitCore, error) {
	const op = "s3-storage.minio-client.minio-client.DownloadFile"

	reader, err := m.Client.GetObject(
		ctx,
		constant.BucketName,
		objectName,
		minio.GetObjectOptions{},
	)
	defer reader.Close()

	if err != nil {
		return models.ImageUnitCore{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.ImageUnitCore{
		Payload: reader,
	}, nil
}

func generateObjectName() string {
	t := time.Now()

	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return fmt.Sprintf(
		"%s.%s",
		formatted,
		"jpg",
	)
}
