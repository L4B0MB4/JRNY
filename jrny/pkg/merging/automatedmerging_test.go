package merging

import (
	"testing"

	"github.com/L4B0MB4/JRNY/jrny/pkg/space"
)

func TestInitialize(t *testing.T) {

	m := SelfConfiguringMerging{}
	if m.knownIdentifiers != nil {
		t.Errorf("Expected knownIdentifierArray to be uninitialized")
	}
	m.Initialize(&space.ResponsibleArea{
		From: 1,
		To:   4,
	})
	if m.knownIdentifiers == nil {
		t.Errorf("Expected knownIdentifierArray to be initialized")
	}

}
