package worker

import (
	"github.com/L4B0MB4/JRNY/pkg/models"
)

type EventPoolWorker interface {
	SetUp()
	Shutdown()
	IsActive() bool
	OnEvent(event *models.Event)
}
