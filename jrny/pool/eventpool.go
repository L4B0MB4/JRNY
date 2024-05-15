package pool

import (
	"context"
	"errors"

	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/rs/zerolog/log"
)

type EventPooler interface {
	Enqueue(event *models.Event) error
	PublishEvents()
}

type EventPool struct {
	queue  chan models.Event
	worker []EventPoolWorker
}

func (e *EventPool) queueInitializedGuard() error {
	if e.queue == nil {
		return errors.New(`queue is not yet initialized`)
	}
	return nil
}
func (e *EventPool) Initialize(factory EventPoolWorkerFactory, ctx context.Context) error {
	if e.queue != nil {
		return errors.New("initialization has been called before")
	}
	e.queue = make(chan models.Event, 500)
	e.worker = factory.Generate()
	go e.PublishEvents()
	go e.onCancel(ctx)
	return nil
}

func (e *EventPool) onCancel(ctx context.Context) {
	err := e.queueInitializedGuard()
	if err != nil {
		return
	}
	<-ctx.Done()
	ctxerr := ctx.Err()
	if ctxerr == context.Canceled {
		log.Debug().Msg("Context was ")
	} else {
		log.Error().Msg("Unexpected context cancelation error")
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
			e.worker[0].OnEvent(&event)
		} else {
			log.Debug().Msg("ending PublishEvents")
			break
		}
	}

}
