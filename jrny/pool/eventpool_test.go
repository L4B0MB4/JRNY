package pool_test

import (
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
	ep := pool.EventPool{}
	ep.Initialize()
	err := ep.Enqueue(&models.Event{})

	if err != nil {
		t.Error("error for initialized eventpool")
	}

}
func TestDoubleInitializedQueue(t *testing.T) {
	ep := pool.EventPool{}
	err := ep.Initialize()
	if err != nil {
		t.Error("error for uninitialized eventpool after initialize")
	}
	err = ep.Initialize()
	if err == nil {
		t.Error("no error for already initialized eventpool")
	}

}
