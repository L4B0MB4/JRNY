package mocks

import "github.com/L4B0MB4/JRNY/jrny/pool"

type TestWorkerFactory struct {
	Worker TestWorker
}

func (factory *TestWorkerFactory) Generate() []pool.EventPoolWorker {
	factory.Worker = TestWorker{}
	workerSlice := []pool.EventPoolWorker{&factory.Worker}
	return workerSlice
}
