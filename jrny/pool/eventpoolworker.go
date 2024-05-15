package pool

import "github.com/L4B0MB4/JRNY/jrny/models"

type EventPoolWorker interface {
	OnEvent(event *models.Event)
}

type LoggingEventPoolWorker struct {
}

func (worker *LoggingEventPoolWorker) OnEvent(event *models.Event) {

}
