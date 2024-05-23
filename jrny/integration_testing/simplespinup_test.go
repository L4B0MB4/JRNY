package integration_testing

import (
	"context"
	"testing"
	"time"

	"github.com/L4B0MB4/JRNY/jrny/integration_testing/mocks"
	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/L4B0MB4/JRNY/jrny/pool"
)

func TestStartsAndRoutesEventsThroughPool(t *testing.T) {
	t.Log("Running TestStartsAndRoutesEventsThroughPool")
	ctx, cancel := context.WithCancel(context.Background())
	factory := &mocks.TestWorkerFactory{}
	eventpool := pool.EventPool{}
	eventpool.Initialize(factory, ctx)
	if factory.Worker.OnEventCalls != 0 {
		t.Error("No calls should've been made")
	}
	eventpool.Enqueue(&models.Event{})
	time.Sleep(100_000)
	if factory.Worker.OnEventCalls != 1 {
		t.Error("One call to eventworker should've been made")
	}
	cancel()
	time.Sleep(50_000_000)
	if factory.Worker.ShutdownCalls != 1 {
		t.Error("One call to shutdown of eventworker should've been made")
	}

}
