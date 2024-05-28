package main

import (
	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/factory"
)

func main() {
	config, workerFactory := setup()
	server.Start(config, workerFactory)
}
func setup() (*configuration.ServerConfiguration, factory.EventPoolWorkerFactory) {

	config := configuration.DefaultConfiguration()
	workerFactory := factory.RabbitMqEventPoolWorkerFactory{
		Config: &config,
	}

	return &config, &workerFactory

}
