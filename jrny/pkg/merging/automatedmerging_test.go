package merging

import (
	"math/big"
	"testing"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/space"
	"github.com/google/uuid"
)

func TestInitialize(t *testing.T) {

	m := SelfConfiguringMerging{}
	if m.knownIdentifiers != nil {
		t.Errorf("Expected knownIdentifierArray to be uninitialized")
	}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(1),
		To:   *big.NewInt(4),
	})
	if m.knownIdentifiers == nil {
		t.Errorf("Expected knownIdentifierArray to be initialized")
	}

}
func TestResponsibleAreaGuard(t *testing.T) {
	m := SelfConfiguringMerging{}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(1),
		To:   *big.NewInt(4),
	})
	myuuid := uuid.MustParse("ed4c5e8f-c512-48ba-b488-bb4be07508e3")
	_, ok := m.knownIdentifiers[myuuid]
	if ok {
		t.Error("Expected item not to be in knownIdentifiers")
	}
	m.Merge(&models.Event{
		Type: "a-type",
		ID:   myuuid,
	})
	_, ok = m.knownIdentifiers[myuuid]
	if ok {
		t.Error("Expected item not to be in knownIdentifiers due to it being out of range")
	}
}

func TestMergingUnkowns(t *testing.T) {

	m := SelfConfiguringMerging{}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(1),
		To:   *big.NewInt(0).SetBytes(uuid.Max[:]),
	})
	myuuid := uuid.MustParse("ed4c5e8f-c512-48ba-b488-bb4be07508e3")
	_, ok := m.knownIdentifiers[myuuid]
	if ok {
		t.Error("Expected item not to be in knownIdentifiers")
	}
	m.Merge(&models.Event{
		Type: "a-type",
		ID:   myuuid,
	})
	_, ok = m.knownIdentifiers[myuuid]
	if !ok {
		t.Error("Expected item to be in knownIdentifiers")
	}

}
