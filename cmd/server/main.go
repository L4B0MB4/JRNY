package main

import (
	"os"

	"github.com/L4B0MB4/JRNY/pkg/configuration"
	"github.com/L4B0MB4/JRNY/pkg/server"
	"github.com/L4B0MB4/JRNY/pkg/server/pool/factory"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, workerFactory := setup()
	server.Start(config, workerFactory)
}
func setup() (*configuration.ServerConfiguration, factory.EventPoolWorkerFactory) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config := configuration.DefaultServerConfiguration()
	workerFactory := factory.RabbitMqEventPoolWorkerFactory{
		Config: &config,
	}

	return &config, &workerFactory

}
