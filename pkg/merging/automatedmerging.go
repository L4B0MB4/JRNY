package merging

import (
	"math/big"

	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/L4B0MB4/JRNY/pkg/space"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type GetOrCreateResponse struct {
	identifier   *models.IdentifierReference
	newlyCreated bool
}

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

	connections := make([]*GetOrCreateResponse, 0)

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

func (s *SelfConfiguringMerging) linkEverything(items []*GetOrCreateResponse) (merges []*models.IdentifierReference) {
	for _, item := range items {
		for _, otherItem := range items {
			if otherItem == item {
				continue
			}
			_, ok := item.identifier.Linked[otherItem.identifier.Self]
			if !ok {
				item.identifier.Linked[otherItem.identifier.Self] = otherItem.identifier
				if !item.newlyCreated {
					merges = append(merges, item.identifier)
				}
			}
		}
	}
	return merges
}

// Returns the IdentifierReference and a boolean indicating if it was newly created
func (s *SelfConfiguringMerging) getOrAddToKnownIdentifiers(id uuid.UUID) *GetOrCreateResponse {
	v, ok := s.knownIdentifiers[id]
	if !ok {
		v = &models.IdentifierReference{
			Self:   id,
			Linked: make(map[[16]byte]*models.IdentifierReference),
		}
		s.knownIdentifiers[id] = v
		return &GetOrCreateResponse{identifier: v, newlyCreated: true}
	}
	return &GetOrCreateResponse{identifier: v, newlyCreated: false}
}
