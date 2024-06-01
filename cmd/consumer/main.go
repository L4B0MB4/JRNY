package main

import (
	"os"

	"github.com/L4B0MB4/JRNY/pkg/configuration"
	"github.com/L4B0MB4/JRNY/pkg/consumer"
	"github.com/L4B0MB4/JRNY/pkg/merging"
	"github.com/L4B0MB4/JRNY/pkg/space"
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

	respArea := space.CreateResponsibleAreas(1)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config := configuration.DefaultConsumerConfiguration()
	merger := merging.SelfConfiguringMerging{}
	merger.Initialize(&respArea[0])
	consumer := consumer.RabbitMqConsumer{
		Config: &config,
		Merger: &merger,
	}

	return &consumer

}
