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
	worker := w.RabbitMqEventPoolWorker{}
	worker.SetUp()
	workerSlice := []w.EventPoolWorker{&worker}
	return workerSlice
}
