package consumer

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMqConsumer struct {
	initialized     bool
	channel         *amqp.Channel
	lifeTimeContext context.Context
	Config          *configuration.ConsumerConfiguration
}

func (c *RabbitMqConsumer) Initialize() error {
	if c.Config == nil {
		err := fmt.Errorf("no config set")
		log.Error().Err(err).Msg("No config set for consumer")
		return err
	}

	conn, err := amqp.Dial(c.Config.QueueConfig.Endpoint)
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to rabbitmq host")
		return err
	}
	c.channel, err = conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return err
	}
	_, err = c.channel.QueueDeclare("events", true, false, false, false, amqp.Table{amqp.QueueTypeArg: amqp.QueueTypeStream,
		amqp.StreamMaxLenBytesArg:         int64(5_000_000_000), // 5 Gb
		amqp.StreamMaxSegmentSizeBytesArg: 500_000_000,          // 500 Mb
		amqp.StreamMaxAgeArg:              "3D",                 // 3 days
	})
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

	c.channel.Qos(2, 0, false)
	msgs, err := c.channel.ConsumeWithContext(c.lifeTimeContext, "events", "", false, false, false, false, amqp.Table{"x-stream-offset": "first"})
	if err != nil {
		log.Error().Err(err).Msg("Error reading message from channel")
		return
	}
	for msg := range msgs {
		readMessage(msg)
	}
}

func readMessage(msg amqp.Delivery) {

	reader := bytes.NewReader(msg.Body)
	decoder := gob.NewDecoder(reader)
	readEvent := models.Event{}
	err := decoder.Decode(&readEvent)
	if err != nil {
		log.Error().Err(err).Bytes("msg", msg.Body).Msg("Error decoding read message")
		return
	}
	log.Debug().Interface("readEvent", readEvent).Msg("Message from queue")
}
