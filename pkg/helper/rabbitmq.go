package helper

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func CreateDefaultQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare("events", true, false, false, false, amqp.Table{amqp.QueueTypeArg: amqp.QueueTypeStream,
		amqp.StreamMaxLenBytesArg:         int64(5_000_000_000), // 5 Gb
		amqp.StreamMaxSegmentSizeBytesArg: 500_000_000,          // 500 Mb
		amqp.StreamMaxAgeArg:              "1h",                 // 3 days
	})
	if err != nil {
		log.Error().Err(err).Msg("Could not declare rabbitmq queue")
		return err
	}
	return nil
}

func CreateAvailabliltyQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare("availablilty", true, false, true, false, nil)
	if err != nil {
		log.Error().Err(err).Msg("Could not declare rabbitmq queue")
		return err
	}
	return nil
}
