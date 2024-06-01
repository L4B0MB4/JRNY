package space

import "github.com/google/uuid"

type ResponsibleArea struct {
	From     int //including
	To       int //excluding
	WorkerId uuid.UUID
}

func CreateResponsibleAreas(amountOfWorkers int) []ResponsibleArea {
	firstSpace, otherSpaces := equalDivideSpace(amountOfWorkers)

	responsibleAreas := make([]ResponsibleArea, amountOfWorkers)

	responsibleAreas[0] = ResponsibleArea{
		From:     0x0,
		To:       firstSpace,
		WorkerId: uuid.New(),
	}
	for i := range amountOfWorkers - 1 {
		responsibleAreas[i+1] = ResponsibleArea{
			From: responsibleAreas[i].To,
			To:   responsibleAreas[i].To + otherSpaces,
		}
	}
	return responsibleAreas

}
