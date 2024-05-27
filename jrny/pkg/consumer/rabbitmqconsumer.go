package consumer

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMqConsumer struct {
	initialized     bool
	channel         *amqp.Channel
	lifeTimeContext context.Context
}

func (c *RabbitMqConsumer) Initialize() error {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to rabbitmq host")
		return err
	}
	c.channel, err = conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return err
	}
	_, err = c.channel.QueueDeclare("events", false, false, false, false, nil)
	if err != nil {
		log.Error().Err(err).Msg("Could not declare rabbitmq queue")
		return err
	}
	c.lifeTimeContext = context.Background() //todo: change to context that is canceled on ctrl-c
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

	msgs, err := c.channel.ConsumeWithContext(c.lifeTimeContext, "events", "", true, false, false, false, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error reading message from channel")
		return
	}
	msg := <-msgs
	reader := bytes.NewReader(msg.Body)
	decoder := gob.NewDecoder(reader)
	readEvent := models.Event{}
	err = decoder.Decode(&readEvent)
	if err != nil {
		log.Error().Err(err).Bytes("msg", msg.Body).Msg("Error decoding read message")
		return
	}
	log.Debug().Interface("readEvent", readEvent).Msg("Message from queue")
}
