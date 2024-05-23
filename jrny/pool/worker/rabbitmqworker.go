package worker

import (
	"bytes"
	"encoding/gob"

	"github.com/L4B0MB4/JRNY/jrny/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMqEventPoolWorker struct {
	queue         *amqp.Queue
	channel       *amqp.Channel
	active        bool
	messageBuffer *bytes.Buffer
	encoder       *gob.Encoder
	connection    *amqp.Connection
	Endpoint      string
}

func (w *RabbitMqEventPoolWorker) SetUp() {
	if w.Endpoint == "" {
		w.Endpoint = "localhost:5672"
	}
	conn, err := amqp.Dial("amqp://guest:guest@" + w.Endpoint + "/")
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to rabbitmq host")
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return
	}
	queue, err := ch.QueueDeclare("events", false, false, false, false, nil)
	if err != nil {
		log.Error().Err(err).Msg("Could not declare rabbitmq queue")
		return
	}

	w.connection = conn
	w.queue = &queue
	w.channel = ch
	w.messageBuffer = &bytes.Buffer{}
	w.encoder = gob.NewEncoder(w.messageBuffer)
	w.active = true

}

func (w *RabbitMqEventPoolWorker) Shutdown() {
	if w.active {

		w.active = false
		w.channel.Close()
		w.connection.Close()
	}
}

func (w *RabbitMqEventPoolWorker) IsActive() bool {

	return w.active

}

func (w *RabbitMqEventPoolWorker) isInitialized() bool {

	if !w.active {
		log.Error().Msg("Trying to use inactive RabbitMq Worker")
		return false
	}
	if w.channel == nil || w.queue == nil {
		w.active = false
		log.Error().Msg("Incorrectly initialized RabbitMq Worker")
		return false
	}
	if w.channel.IsClosed() {
		w.active = false
		log.Error().Msg("RabbitMq Worker's channel is already closed")
		return false
	}
	return true

}

func (w *RabbitMqEventPoolWorker) OnEvent(event *models.Event) {
	if !w.isInitialized() {
		return
	}
	w.messageBuffer.Reset()
	err := w.encoder.Encode(event)
	if err != nil {
		log.Error().Err(err).Msg("Could not encode event into binary data")
		return
	}
	err = w.channel.Publish("", w.queue.Name, false, false, amqp.Publishing{
		ContentType: "application/octet-stream",
		Body:        w.messageBuffer.Bytes(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Could not publish event")
		return
	}

}
