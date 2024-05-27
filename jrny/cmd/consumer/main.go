package main

import (
	"os"

	"github.com/L4B0MB4/JRNY/jrny/pkg/consumer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Debug().Msg("Launching consumer")
	defer log.Debug().Msg("Stopped consumer")
	var c consumer.Consumer = &consumer.RabbitMqConsumer{}

	c.Initialize()

}
