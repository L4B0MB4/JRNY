package pool

type EventPoolWorkerFactory interface {
	Generate() []EventPoolWorker
}

type DefaultEventPoolWorkerFactory struct {
}

func (factory *DefaultEventPoolWorkerFactory) Generate() []EventPoolWorker {
	workerSlice := []EventPoolWorker{&LoggingEventPoolWorker{}}
	return workerSlice
}
