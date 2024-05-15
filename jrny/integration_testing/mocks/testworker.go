package mocks

import "github.com/L4B0MB4/JRNY/jrny/models"

type TestWorker struct {
	Calls int
}

func (worker *TestWorker) OnEvent(event *models.Event) {
	worker.Calls++
}
