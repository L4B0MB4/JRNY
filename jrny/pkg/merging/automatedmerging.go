package merging

import (
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
	s.knownIdentifiers = make(map[[16]byte]*models.IdentifierReference, r.To-r.From)
}

func (s *SelfConfiguringMerging) Merge(event *models.Event) {
	id := event.ID
	v, ok := s.knownIdentifiers[id]
	if !ok {
		v = &models.IdentifierReference{
			Self: id,
		}
	}

	log.Debug().Any("v", v).Msg("Logging v")

}
