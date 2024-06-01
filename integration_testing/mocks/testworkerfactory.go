package mocks

import (
	"github.com/L4B0MB4/JRNY/pkg/server/pool/worker"
)

type TestWorkerFactory struct {
	Worker TestWorker
}

func (factory *TestWorkerFactory) Generate() []worker.EventPoolWorker {
	factory.Worker = TestWorker{}
	factory.Worker.SetUp()
	workerSlice := []worker.EventPoolWorker{&factory.Worker}
	return workerSlice
}
