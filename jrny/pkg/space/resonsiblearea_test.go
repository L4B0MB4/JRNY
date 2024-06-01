package space

import "testing"

func TestCreateResponsibleAreas(t *testing.T) {

	responsibleAreaWorkers := CreateResponsibleAreas(4)

	if len(responsibleAreaWorkers) != 4 {
		t.Errorf("Size of areas should be 4 but is %v", len(responsibleAreaWorkers))
		return
	}

	if responsibleAreaWorkers[0].From != 0x0 || responsibleAreaWorkers[0].To != 0x40000000 {
		t.Errorf("Responsible area has not the right size: %x - %x", responsibleAreaWorkers[0].From, responsibleAreaWorkers[0].To)
	}

	if responsibleAreaWorkers[2].From != 0x40000000*2 || responsibleAreaWorkers[2].To != 0x40000000*3 {
		t.Errorf("Responsible area has not the right size: %x - %x", responsibleAreaWorkers[2].From, responsibleAreaWorkers[2].To)
	}

	if responsibleAreaWorkers[3].From != 0x40000000*3 || responsibleAreaWorkers[3].To != 0x100000000 {
		t.Errorf("Responsible area has not the right size: %x - %x", responsibleAreaWorkers[3].From, responsibleAreaWorkers[3].To)
	}
}
