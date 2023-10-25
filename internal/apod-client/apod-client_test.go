package apod_client

import (
	"fmt"
	"github.com/degeboman/betera-test-task/internal/config"
	minio_client "github.com/degeboman/betera-test-task/internal/s3-storage/minio-client"
	"testing"
)

func TestApodClient_DownloadImage(t *testing.T) {

	mockAc := ApodClient{
		s3storage: minio_client.New(config.Config{
			MinioAuthData: config.MinioAuthData{
				Password: "minio123",
				Url:      "localhost:9000",
				User:     "minio",
				Ssl:      false,
			},
		}),
	}

	err := mockAc.s3storage.Connect()

	if err != nil {
		fmt.Println(err)
	}

	const url = "https://apod.nasa.gov/apod/image/2203/WhaleAurora_Strand_960.jpg"

	mockAc.DownloadImage(url)
}
