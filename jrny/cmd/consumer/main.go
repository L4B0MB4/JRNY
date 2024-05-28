package main

import (
	"os"

	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	"github.com/L4B0MB4/JRNY/jrny/pkg/consumer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Debug().Msg("Launching consumer")
	c := setup()
	defer log.Debug().Msg("Stopped consumer")
	err := c.Initialize()
	if err != nil {
		log.Error().Err(err).Msg("Stopping consumer, error during initialization")
		return
	}
	c.Consume()

}
func setup() consumer.Consumer {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config := configuration.DefaultConsumerConfiguration()
	consumer := consumer.RabbitMqConsumer{
		Config: &config,
	}
	return &consumer

}
