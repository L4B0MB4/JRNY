package worker

import (
	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/rs/zerolog/log"
)

type LoggingEventPoolWorker struct {
}

func (worker *LoggingEventPoolWorker) OnEvent(event *models.Event) {
	log.Debug().Interface("event", event).Msg("")
}
