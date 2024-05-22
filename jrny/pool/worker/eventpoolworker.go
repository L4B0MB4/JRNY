package worker

import (
	"github.com/L4B0MB4/JRNY/jrny/models"
)

type EventPoolWorker interface {
	OnEvent(event *models.Event)
}
