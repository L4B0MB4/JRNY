package mocks

import "github.com/L4B0MB4/JRNY/jrny/pkg/models"

type TestWorker struct {
	OnEventCalls  int
	ShutdownCalls int
}

func (worker *TestWorker) OnEvent(event *models.Event) {
	worker.OnEventCalls++
}
func (worker *TestWorker) IsActive() bool {
	return true

}
func (worker *TestWorker) SetUp() {
}
func (worker *TestWorker) Shutdown() {
	worker.ShutdownCalls++
}
