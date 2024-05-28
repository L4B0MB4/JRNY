package factory

import (
	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	w "github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/worker"
)

type RabbitMqEventPoolWorkerFactory struct {
	Config *configuration.ServerConfiguration
}

func (factory *RabbitMqEventPoolWorkerFactory) Generate() []w.EventPoolWorker {
	worker := &w.RabbitMqEventPoolWorker{
		Endpoint: factory.Config.QueueConfig.Endpoint,
	}
	worker.SetUp()
	workerSlice := []w.EventPoolWorker{worker}
	return workerSlice
}
