package factory

import (
	w "github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/worker"
)

type EventPoolWorkerFactory interface {
	Generate() []w.EventPoolWorker
}