package merging

import (
	"math/big"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/space"
	"github.com/rs/zerolog/log"
)

type SelfConfiguringMerging struct {
	responsibleArea  *space.ResponsibleArea
	knownIdentifiers map[[16]byte]*models.IdentifierReference
}

func (s *SelfConfiguringMerging) SelfConfigure() {
	//todo self configure

}

func (s *SelfConfiguringMerging) Initialize(r *space.ResponsibleArea) {
	s.responsibleArea = r
	s.knownIdentifiers = make(map[[16]byte]*models.IdentifierReference, 1000)
}

func (s *SelfConfiguringMerging) ResponsibleAreaGuard(id [16]byte) bool {
	idBigInt := big.NewInt(0).SetBytes(id[:])
	if idBigInt.Cmp(&s.responsibleArea.From) >= 0 && idBigInt.Cmp(&s.responsibleArea.To) == -1 {
		return true
	}
	return false
}

func (s *SelfConfiguringMerging) Merge(event *models.Event) {
	id := event.ID
	if !s.ResponsibleAreaGuard(id) {
		return
	}

	v, ok := s.knownIdentifiers[id]
	if !ok {
		v = &models.IdentifierReference{
			Self: id,
		}
		s.knownIdentifiers[id] = v
	}

	log.Debug().Any("v", v).Msg("Logging v")

}
