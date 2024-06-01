package space

import (
	"math/big"

	"github.com/google/uuid"
)

type ResponsibleArea struct {
	From     big.Int //including
	To       big.Int //excluding
	WorkerId uuid.UUID
}

func CreateResponsibleAreas(amountOfWorkers int64) []ResponsibleArea {
	responsibleAreas := make([]ResponsibleArea, amountOfWorkers)
	firstSpace, otherSpaces := equalDivideSpace(amountOfWorkers)

	responsibleAreas[0] = ResponsibleArea{
		From:     *big.NewInt(0),
		To:       *firstSpace,
		WorkerId: uuid.New(),
	}
	for i := range amountOfWorkers - 1 {
		to := big.NewInt(0)
		to.Add(&responsibleAreas[i].To, otherSpaces)

		responsibleAreas[i+1] = ResponsibleArea{
			From: responsibleAreas[i].To,
			To:   *to,
		}
	}
	return responsibleAreas

}
