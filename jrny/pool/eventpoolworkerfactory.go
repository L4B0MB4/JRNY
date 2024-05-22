package pool

import (
	w "github.com/L4B0MB4/JRNY/jrny/pool/worker"
)

type EventPoolWorkerFactory interface {
	Generate() []w.EventPoolWorker
}

type DefaultEventPoolWorkerFactory struct {
	useLoggingWorker bool
}

func (factory *DefaultEventPoolWorkerFactory) Generate() []w.EventPoolWorker {
	var worker w.EventPoolWorker
	if !factory.useLoggingWorker {
		worker = &w.RabbitMqEventPoolWorker{}
	} else {
		worker = &w.LoggingEventPoolWorker{}
	}
	worker.SetUp()
	workerSlice := []w.EventPoolWorker{worker}
	return workerSlice
}

func (factory *DefaultEventPoolWorkerFactory) UseLoggingWorker() {
	factory.useLoggingWorker = true
}
