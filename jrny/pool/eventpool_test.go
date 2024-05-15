package pool_test

import (
	"context"
	"testing"

	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/L4B0MB4/JRNY/jrny/pool"
)

func TestUnInitializedQueue(t *testing.T) {
	ep := pool.EventPool{}
	err := ep.Enqueue(&models.Event{})
	if err == nil {
		t.Error("no error for uninitialized eventpool")
	}

}

func TestInitializedQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ep := pool.EventPool{}
	factory := &pool.DefaultEventPoolWorkerFactory{}
	ep.Initialize(factory, ctx)
	err := ep.Enqueue(&models.Event{})

	if err != nil {
		t.Error("error for initialized eventpool")
	}

}
func TestDoubleInitializedQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	factory := &pool.DefaultEventPoolWorkerFactory{}
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
