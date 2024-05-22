package pool

import (
	w "github.com/L4B0MB4/JRNY/jrny/pool/worker"
)

type EventPoolWorkerFactory interface {
	Generate() []w.EventPoolWorker
}

type DefaultEventPoolWorkerFactory struct {
}

func (factory *DefaultEventPoolWorkerFactory) Generate() []w.EventPoolWorker {
	workerSlice := []w.EventPoolWorker{&w.LoggingEventPoolWorker{}}
	return workerSlice
}
