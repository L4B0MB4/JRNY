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

func TestLinkingWithinMerge(t *testing.T) {

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
	merges := m.Merge(&models.Event{
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
	if len(merges) != 1 || len(merges[0]) != 3 {
		t.Errorf("There should be only one merge containing 3 items")
		return
	}

	_, ok := merges[0][myuuid1]
	if !ok {
		t.Errorf("myuuid1 should be present in merge")
		return
	}

	_, ok = merges[0][myuuid2]
	if !ok {
		t.Errorf("myuuid2 should be present in merge")
		return
	}

	_, ok = merges[0][myuuid7]
	if !ok {
		t.Errorf("myuuid7 should be present in merge")
		return
	}

	myuuid8 := uuid.MustParse("00000000-0000-0000-0000-000000000007")
	merges = m.Merge(&models.Event{
		Type: "x-type",
		ID:   myuuid2,
		Relationships: map[string][]models.Relation{
			"rel-type": {
				{
					Type: "j-type",
					ID:   myuuid8,
				},
			},
		},
	})
	if len(merges) != 1 || len(merges[0]) != 2 {
		t.Errorf("There should be only one merge containing 2 items in second iteration")
		return
	}

	_, ok = merges[0][myuuid2]
	if !ok {
		t.Errorf("myuuid2 should be present in second merge")
		return
	}

	_, ok = merges[0][myuuid8]
	if !ok {
		t.Errorf("myuuid8 should be present in second merge")
		return
	}

}

func TestOneSidedMerge(t *testing.T) {
	mBig := SelfConfiguringMerging{}
	mBig.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(4),
		To:   *big.NewInt(8),
	})
	mSmall := SelfConfiguringMerging{}
	mSmall.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(0),
		To:   *big.NewInt(4),
	})

	myuuid1 := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	myuuid7 := uuid.MustParse("00000000-0000-0000-0000-000000000007")

	f := &models.Event{
		Type: "a",
		ID:   myuuid1,
		Relationships: map[string][]models.Relation{
			"b": {
				models.Relation{
					Type: "c",
					ID:   myuuid7,
				},
			},
		},
	}

	merges := mSmall.Merge(f)

	if len(merges) != 1 {
		t.Errorf("Merges should contain 1 item but contain %v", len(merges))
		return
	}

	merges = mBig.Merge(f)

	if len(merges) != 0 {
		t.Errorf("Merges should contain 0 items but contain %v", len(merges))
		return
	}
}

func TestOneSidedMergeWithAdditionalData(t *testing.T) {
	mBig := SelfConfiguringMerging{}
	mBig.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(4),
		To:   *big.NewInt(8),
	})
	mSmall := SelfConfiguringMerging{}
	mSmall.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(0),
		To:   *big.NewInt(4),
	})

	myuuid1 := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	myuuid2 := uuid.MustParse("00000000-0000-0000-0000-000000000002")
	myuuid7 := uuid.MustParse("00000000-0000-0000-0000-000000000007")

	first := &models.Event{
		Type: "a",
		ID:   myuuid1,
		Relationships: map[string][]models.Relation{
			"b": {
				models.Relation{
					Type: "c",
					ID:   myuuid7,
				},
			},
		},
	}
	mSmall.Merge(first)
	mBig.Merge(first)

	second := &models.Event{
		Type: "a",
		ID:   myuuid7,
		Relationships: map[string][]models.Relation{
			"b": {
				models.Relation{
					Type: "c",
					ID:   myuuid2,
				},
			},
		},
	}
	merges := mSmall.Merge(second)
	if len(merges) != 0 {
		t.Errorf("Merges should contain 0 items but contain %v", len(merges))
		return
	}

	merges = mBig.Merge(second)

	if len(merges) != 1 {
		t.Errorf("Merges should contain 1 item but contain %v", len(merges))
		return
	}

	_, ok := merges[0][myuuid7]
	if !ok {
		t.Error("myuuid7 should be in the merge list")
		return
	}
	_, ok = merges[0][myuuid2]
	if !ok {
		t.Error("myuuid2 should be in the merge list")
		return
	}
}

func TestOneSidedMergeForThreeSeparateAreas(t *testing.T) {
	mBig := SelfConfiguringMerging{}
	mBig.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(8),
		To:   *big.NewInt(99),
	})
	mMid := SelfConfiguringMerging{}
	mMid.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(4),
		To:   *big.NewInt(8),
	})
	mSmall := SelfConfiguringMerging{}
	mSmall.Initialize(&space.ResponsibleArea{
		From: *big.NewInt(0),
		To:   *big.NewInt(4),
	})

	myuuid1 := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	myuuid7 := uuid.MustParse("00000000-0000-0000-0000-000000000007")
	myuuid15 := uuid.MustParse("00000000-0000-0000-0000-00000000000F")

	f := &models.Event{
		Type: "a",
		ID:   myuuid15,
		Relationships: map[string][]models.Relation{
			"b": {
				models.Relation{
					Type: "c",
					ID:   myuuid7,
				},

				models.Relation{
					Type: "d",
					ID:   myuuid1,
				},
			},
		},
	}

	merges := mSmall.Merge(f)

	if len(merges) != 0 {
		t.Errorf("Merges should contain 0 items but contain %v", len(merges))
		return
	}

	merges = mMid.Merge(f)

	if len(merges) != 0 {
		t.Errorf("Merges should contain 0 items but contain %v", len(merges))
		return
	}

	merges = mBig.Merge(f)

	if len(merges) != 1 {
		t.Errorf("Merges should contain 1 item but contain %v", len(merges))
		return
	}
}
