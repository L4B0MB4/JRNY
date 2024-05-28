package factory

import (
	w "github.com/L4B0MB4/JRNY/jrny/pkg/pool/worker"
)

type LoggingEventPoolWorkerFactory struct {
}

func (factory *LoggingEventPoolWorkerFactory) Generate() []w.EventPoolWorker {
	worker := &w.LoggingEventPoolWorker{}
	worker.SetUp()
	workerSlice := []w.EventPoolWorker{worker}
	return workerSlice
}
