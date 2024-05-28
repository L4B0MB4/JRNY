// Package contains pool that activates workers and sends them messages
// on receive over a channel.

package pool

import (
	"context"
	"errors"
	"time"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/factory"
	w "github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/worker"
	"github.com/rs/zerolog/log"
)

type EventPooler interface {
	Enqueue(event *models.Event) error
	PublishEvents()
}

type EventPool struct {
	queue  chan models.Event
	worker []w.EventPoolWorker
}

func (e *EventPool) queueInitializedGuard() error {
	if e.queue == nil {
		return errors.New(`queue is not yet initialized`)
	}
	return nil
}
func (e *EventPool) Initialize(factory factory.EventPoolWorkerFactory, ctx context.Context) error {
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
	log.Debug().Msg(time.Now().String())
	ctxerr := ctx.Err()
	if ctxerr == context.Canceled {
		log.Debug().Msg("Context was canceled")
	} else {
		log.Error().Msg("Unexpected context cancelation error")
	}
	close(e.queue)
	for _, w := range e.worker {
		w.Shutdown()
	}
	log.Debug().Msg(time.Now().String())
}

/*
Takes validated event models and enqueues them into the pool
*/
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
