package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMqConsumer struct {
	initialized bool
}

func (c *RabbitMqConsumer) Initialize() error {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to rabbitmq host")
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return err
	}
	_, err = ch.QueueDeclare("events", false, false, false, false, nil)
	if err != nil {
		log.Error().Err(err).Msg("Could not declare rabbitmq queue")
		return err
	}
	c.initialized = true
	return nil

}

func (c *RabbitMqConsumer) initializedGuard() bool {
	return c.initialized

}

func (c *RabbitMqConsumer) Consume() {

	if !c.initializedGuard() {
		return
	}

}
