package pool

import (
	"errors"

	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/rs/zerolog/log"
)

type EventPooler interface {
	Enqueue(event *models.Event) error
	PublishEvents()
}

type EventPool struct {
	queue chan models.Event
}

func (e *EventPool) queueInitializedGuard() error {
	if e.queue == nil {
		return errors.New(`queue is not yet initialized`)
	}
	return nil
}
func (e *EventPool) Initialize() error {
	if e.queue != nil {
		return errors.New("initialization has been called before")
	}
	e.queue = make(chan models.Event, 500)
	go e.PublishEvents()
	return nil
}
func (e *EventPool) Shutdown() {
	err := e.queueInitializedGuard()
	if err != nil {
		return
	}
	close(e.queue)
}

// Takes validated event models and enqueues them into the pool
func (e *EventPool) Enqueue(event *models.Event) error {
	err := e.queueInitializedGuard()
	if err != nil {
		return err
	}
	e.queue <- *event
	return nil
}

func (e *EventPool) PublishEvents() {
	var ok bool = true
	var event models.Event
	for {
		event, ok = <-e.queue
		if ok {
			log.Debug().Interface("event", event).Msg("hello")
		} else {
			log.Debug().Msg("quit")
			break
		}
	}

}
