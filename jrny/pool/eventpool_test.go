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
		t.Error("no error for unitialized eventpool")
	}

}

func TestInitializedQueue(t *testing.T) {
	ep := pool.EventPool{}
	ep.Initialize()
	err := ep.Enqueue(&models.Event{})

	if err != nil {
		t.Error("error for initialzed eventpool")
	}

}

func FailingTest(t *testing.T) {
	t.Error("failing test")
}
