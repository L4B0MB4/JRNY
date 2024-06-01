package worker

import (
	"bytes"
	"encoding/gob"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
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
	conn, err := amqp.Dial(w.Endpoint)
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to rabbitmq host")
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect open rabbitmq channel")
		return
	}
	queue, err := ch.QueueDeclare("events", true, false, false, false, amqp.Table{amqp.QueueTypeArg: amqp.QueueTypeStream,
		amqp.StreamMaxLenBytesArg:         int64(5_000_000_000), // 5 Gb
		amqp.StreamMaxSegmentSizeBytesArg: 500_000_000,          // 500 Mb
		amqp.StreamMaxAgeArg:              "1h",
	})
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
	w.encoder = gob.NewEncoder(w.messageBuffer)
	//todo: encoders apparently work by sending type once and then only values. that's too dangerous in this case
	//as we do not know who has already read the type and who hasn't. Therefore this is a fix to be removed later
	// when performance matters
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
