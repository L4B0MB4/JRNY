package pool

import (
	w "github.com/L4B0MB4/JRNY/jrny/pool/worker"
)

type EventPoolWorkerFactory interface {
	Generate() []w.EventPoolWorker
}

type DefaultEventPoolWorkerFactory struct {
	useQueueWorker bool
	queueEndpoint  string
}

func (factory *DefaultEventPoolWorkerFactory) Generate() []w.EventPoolWorker {
	var worker w.EventPoolWorker
	if factory.useQueueWorker {
		queueWorker := w.RabbitMqEventPoolWorker{}
		queueWorker.Endpoint = factory.queueEndpoint
		worker = &queueWorker

	} else {
		worker = &w.LoggingEventPoolWorker{}
	}
	worker.SetUp()
	workerSlice := []w.EventPoolWorker{worker}
	return workerSlice
}

func (factory *DefaultEventPoolWorkerFactory) UseQueueWorker(queueEndpoint string) {
	factory.useQueueWorker = true
	factory.queueEndpoint = queueEndpoint
}
