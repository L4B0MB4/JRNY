package merging

import (
	"math/big"
	"slices"
	"testing"

	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/L4B0MB4/JRNY/pkg/space"
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
	myuuid = uuid.MustParse("00000000-0000-0000-0000-000000000004")
	m.Merge(&models.Event{
		Type: "b-type",
		ID:   myuuid,
	})
	_, ok = m.knownIdentifiers[myuuid]
	if ok {
		t.Error("Expected item not to be in knownIdentifiers due to it being out of range")
	}
}
func TestResponsibleAreaGuardFitsArea(t *testing.T) {
	m := SelfConfiguringMerging{}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(1),
		To:   *big.NewInt(4),
	})
	myuuid := uuid.MustParse("00000000-0000-0000-0000-000000000003")
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
		t.Error("Expected item to be in knownIdentifiers due to it being in range")
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

/*
func TestMergingEndResult(t *testing.T) {

	m := SelfConfiguringMerging{}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(1),
		To:   *big.NewInt(0).SetBytes(uuid.Max[:]),
	})
	relUuid := uuid.MustParse("00000000-0000-0000-0000-0000000000FF")
	myuuid1 := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	item1 := &models.Event{
		Type: "a-type",
		ID:   myuuid1,
		Relationships: map[string][]models.Relation{
			"my-key": {
				{Type: "my-rel-type", ID: relUuid},
			},
		},
	}
	myuuid2 := uuid.MustParse("00000000-0000-0000-0000-000000000002")
	item2 := &models.Event{
		Type: "a-type",
		ID:   myuuid2,
		Relationships: map[string][]models.Relation{
			"my-key": {
				{Type: "my-rel-type", ID: relUuid},
			},
		},
	}
	myuuid3 := uuid.MustParse("00000000-0000-0000-0000-000000000003")
	item3 := &models.Event{
		Type: "a-type",
		ID:   myuuid3,
		Relationships: map[string][]models.Relation{
			"my-key": {
				{Type: "my-rel-type", ID: relUuid},
			},
		},
	}

	m.Merge(item1)
	m.Merge(item2)
	m.Merge(item3)

	if len(m.knownIdentifiers) != 4 {
		t.Errorf("Should contain 4 items 3 normal ones and one relationship item but contains %v", len(m.knownIdentifiers))
	}

	v, ok := m.knownIdentifiers[relUuid]
	if !ok {
		t.Error("Should have relationship item as kown identifier")
	}

	if len(v.Linked) != 3 {
		t.Errorf("Should contain 3 itemsbut contains %v", len(v.Linked))
	}
	_, ok = v.Linked[myuuid1]
	if !ok {
		t.Error("Should contain first identifier")
	}
	_, ok = v.Linked[myuuid2]
	if !ok {
		t.Error("Should contain second identifier")
	}
	_, ok = v.Linked[myuuid3]
	if !ok {
		t.Error("Should contain thrid identifier")
	}

}
*/

func TestMerge(t *testing.T) {

	m := SelfConfiguringMerging{}
	m.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(0),
		To:   *big.NewInt(4),
	})

	myuuid1 := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	myuuid2 := uuid.MustParse("00000000-0000-0000-0000-000000000002")
	myuuid7 := uuid.MustParse("00000000-0000-0000-0000-000000000007")
	if len(m.knownIdentifiers) != 0 {
		t.Error("KnownIdentifiers should be empty")
		return
	}
	m.Merge(&models.Event{
		Type: "b-type",
		ID:   myuuid1,
		Relationships: map[string][]models.Relation{
			"rel-type": {
				{
					Type: "some-type",
					ID:   myuuid2,
				},
				{
					Type: "some-type",
					ID:   myuuid7,
				},
			},
		},
	})
	if len(m.knownIdentifiers) != 2 {
		t.Errorf("KnownIdentifiers should be contain 2 elements but contains %v", len(m.knownIdentifiers))
		return
	}

	v := m.knownIdentifiers[myuuid1]

	if v == nil || len(v.Linked) != 2 || !slices.Contains(v.Linked, [16]byte(myuuid2)) || !slices.Contains(v.Linked, [16]byte(myuuid7)) {
		t.Error("Incorrectly connected identifier")
		return
	}

	v = m.knownIdentifiers[myuuid2]

	if v == nil || len(v.Linked) != 2 || !slices.Contains(v.Linked, [16]byte(myuuid1)) || !slices.Contains(v.Linked, [16]byte(myuuid7)) {
		t.Error("Incorrectly connected identifier")
		return
	}

}
