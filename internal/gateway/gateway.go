package gateway

import (
	s3storage "github.com/degeboman/betera-test-task/internal/s3-storage"
	minioclient "github.com/degeboman/betera-test-task/internal/s3-storage/minio-client"
	"github.com/degeboman/betera-test-task/internal/storage"
	postgresclient "github.com/degeboman/betera-test-task/internal/storage/postgres-client"
)

type Gateway struct {
	storage.ApodStorage
	s3storage.S3Storage
}

func SetupGateway(pc postgresclient.PostgresClient, mc *minioclient.MinioClient) Gateway {
	return Gateway{
		ApodStorage: pc,
		S3Storage:   mc,
	}
}
