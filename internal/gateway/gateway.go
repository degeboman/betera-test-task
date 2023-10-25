package gateway

import (
	"github.com/degeboman/betera-test-task/internal/storage"
	postgresclient "github.com/degeboman/betera-test-task/internal/storage/postgres-client"
	"go.uber.org/fx"
)

type Gateway struct {
	fx.Out
	apodStorage storage.ApodStorage
}

func SetupGateway(pc postgresclient.PostgresClient) Gateway {
	return Gateway{
		apodStorage: storage.ApodStorageImpl{Pc: pc},
	}
}
