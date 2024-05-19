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
	ctx, cancel := context.WithCancel(context.Background())
	factory := &mocks.TestWorkerFactory{}
	defer cancel()
	eventpool := pool.EventPool{}
	eventpool.Initialize(factory, ctx)
	if factory.Worker.Calls != 0 {
		t.Error("No calls should've been made")
	}
	eventpool.Enqueue(&models.Event{})
	time.Sleep(1_000)
	if factory.Worker.Calls != 1 {
		t.Error("One call should've been made")
	}

}
