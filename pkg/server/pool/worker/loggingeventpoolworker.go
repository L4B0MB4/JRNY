package worker

import (
	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/rs/zerolog/log"
)

type LoggingEventPoolWorker struct {
}

func (worker *LoggingEventPoolWorker) OnEvent(event *models.Event) {
	log.Debug().Interface("event", event).Msg("")
}

func (worker *LoggingEventPoolWorker) IsActive() bool {
	return true

}

func (worker *LoggingEventPoolWorker) SetUp() {
}

func (worker *LoggingEventPoolWorker) Shutdown() {
}
