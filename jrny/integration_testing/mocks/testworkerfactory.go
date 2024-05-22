package mocks

import (
	"github.com/L4B0MB4/JRNY/jrny/pool/worker"
)

type TestWorkerFactory struct {
	Worker TestWorker
}

func (factory *TestWorkerFactory) Generate() []worker.EventPoolWorker {
	factory.Worker = TestWorker{}
	workerSlice := []worker.EventPoolWorker{&factory.Worker}
	return workerSlice
}
