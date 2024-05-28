package pool_test

import (
	"context"
	"testing"

	"github.com/L4B0MB4/JRNY/jrny/pkg/configuration"
	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool"
	"github.com/L4B0MB4/JRNY/jrny/pkg/server/pool/factory"
)

func TestUnInitializedQueue(t *testing.T) {
	t.Log("Running TestUnInitializedQueue")
	ep := pool.EventPool{}
	err := ep.Enqueue(&models.Event{})
	if err == nil {
		t.Error("no error for uninitialized eventpool")
	}

}

func TestInitializedQueue(t *testing.T) {
	t.Log("Running TestInitializedQueue")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ep := pool.EventPool{}
	config := configuration.DefaultServerConfiguration()
	factory := &factory.RabbitMqEventPoolWorkerFactory{
		Config: &config,
	}
	ep.Initialize(factory, ctx)
	err := ep.Enqueue(&models.Event{})

	if err != nil {
		t.Error("error for initialized eventpool")
	}

}
func TestDoubleInitializedQueue(t *testing.T) {
	t.Log("Running TestDoubleInitializedQueue")
	ctx, cancel := context.WithCancel(context.Background())
	config := configuration.DefaultServerConfiguration()
	factory := &factory.RabbitMqEventPoolWorkerFactory{
		Config: &config,
	}
	defer cancel()
	ep := pool.EventPool{}
	err := ep.Initialize(factory, ctx)
	if err != nil {
		t.Error("error for uninitialized eventpool after initialize")
	}
	err = ep.Initialize(factory, ctx)
	if err == nil {
		t.Error("no error for already initialized eventpool")
	}

}
