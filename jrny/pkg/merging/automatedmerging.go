package merging

import (
	"math/big"

	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/space"
	"github.com/google/uuid"
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

	connections := make([]*models.IdentifierReference, 0)

	v := s.getOrAddToKnownIdentifiers(id)
	connections = append(connections, v)

	for _, relArr := range event.Relationships {
		for _, relationship := range relArr {
			v = s.getOrAddToKnownIdentifiers(relationship.ID)
			connections = append(connections, v)
		}
	}

	s.linkEverything(connections)

	log.Debug().Any("event", event).Msg("Added event to knownidentifiers")

}

func (s *SelfConfiguringMerging) linkEverything(items []*models.IdentifierReference) {
	for _, item := range items {
		for _, otherItem := range items {
			if otherItem == item {
				continue
			}
			_, ok := item.Linked[otherItem.Self]
			if !ok {
				item.Linked[otherItem.Self] = otherItem
			}
		}
	}
}

func (s *SelfConfiguringMerging) getOrAddToKnownIdentifiers(id uuid.UUID) *models.IdentifierReference {
	v, ok := s.knownIdentifiers[id]
	if !ok {
		v = &models.IdentifierReference{
			Self:   id,
			Linked: make(map[[16]byte]*models.IdentifierReference),
		}
		s.knownIdentifiers[id] = v
	}
	return v
}
