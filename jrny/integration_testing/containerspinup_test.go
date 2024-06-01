package integration_testing

import (
	"bytes"
	"context"
	"encoding/gob"
	"testing"
	"time"

	rabbitmq "github.com/L4B0MB4/JRNY/jrny/integration_testing/rabbitmq"
	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	"github.com/L4B0MB4/JRNY/jrny/pkg/helper"
	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/factory"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func TestIntegratesWithRabbitMq(t *testing.T) {
	t.Log("Running TestStartsAndRoutesEventsThroughPool")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	container, err := rabbitmq.SetUpContainer(ctx)
	if err != nil {
		t.Errorf("Error starting container %v", err)
		return
	}
	defer container.Terminate(ctx)

	endpoint, err := container.PortEndpoint(context.Background(), "5672/tcp", "")
	if err != nil {
		t.Error(err)
		return
	}

	channel, err := createRabbitMqReader(endpoint)
	if err != nil {
		t.Error(err)
		return
	}

	config := configuration.DefaultServerConfiguration()
	config.QueueConfig.Endpoint = "amqp://guest:guest@" + endpoint + "/"
	factory := &factory.RabbitMqEventPoolWorkerFactory{
		Config: &config,
	}
	eventpool := pool.EventPool{}
	err = eventpool.Initialize(factory, ctx)
	if err != nil {
		t.Error(err)
		return
	}
	err = eventpool.Enqueue(&models.Event{Type: "mytype"})
	if err != nil {
		t.Error(err)
		return
	}
	timeCtx, cancelTimeCtx := context.WithTimeout(context.Background(), 100*time.Millisecond*1000)
	defer cancelTimeCtx()
	err = channel.Qos(2, 0, false)
	if err != nil {
		t.Error(err)
		return
	}
	msgs, err := channel.ConsumeWithContext(timeCtx, "events", "", false, false, false, false, amqp.Table{"x-stream-offset": "first"})
	if err != nil {
		t.Error(err)
		return
	}
	msg := <-msgs
	msg.Ack(true)
	reader := bytes.NewReader(msg.Body)
	var decoder = gob.NewDecoder(reader)
	readEvent := models.Event{}
	err = decoder.Decode(&readEvent)
	if err != nil {
		t.Error(err)
		return
	}
	log.Debug().Interface("readEvent", readEvent).Msg("Message from queue")

	if readEvent.Type != "mytype" {
		t.Error("Type of read event should be 'mytype'")
	}

}

func createRabbitMqReader(endpoint string) (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@" + endpoint + "/")
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to rabbitmq host")
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return nil, err
	}
	err = helper.CreateDefaultQueue(ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
