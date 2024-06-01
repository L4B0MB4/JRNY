package consumer

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/L4B0MB4/JRNY/pkg/configuration"
	"github.com/L4B0MB4/JRNY/pkg/helper"
	"github.com/L4B0MB4/JRNY/pkg/merging"
	"github.com/L4B0MB4/JRNY/pkg/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMqConsumer struct {
	initialized     bool
	channel         *amqp.Channel
	lifeTimeContext context.Context
	Config          *configuration.ConsumerConfiguration
	Merger          merging.Merger
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
	err = helper.CreateDefaultQueue(c.channel)
	if err != nil {
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
		c.readMessage(msg)
		msg.Ack(true)
	}
}

func (c *RabbitMqConsumer) readMessage(msg amqp.Delivery) {

	reader := bytes.NewReader(msg.Body)
	decoder := gob.NewDecoder(reader)
	readEvent := models.Event{}
	err := decoder.Decode(&readEvent)
	if err != nil {
		log.Error().Err(err).Bytes("msg", msg.Body).Msg("Error decoding read message")
		return
	}
	log.Debug().Interface("readEvent", readEvent).Msg("Message from queue")

	c.Merger.Merge(&readEvent)
	log.Debug().Str("millisecondsToProcess", time.Now().UTC().Sub(readEvent.ReceivedAt).String()).Msg("Finished working on event")
}
