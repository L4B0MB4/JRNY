package merging

import (
	"math/big"

	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/L4B0MB4/JRNY/pkg/space"
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

func (s *SelfConfiguringMerging) responsibleAreaGuard(id [16]byte) bool {
	idBigInt := big.NewInt(0).SetBytes(id[:])
	if idBigInt.Cmp(&s.responsibleArea.From) >= 0 && idBigInt.Cmp(&s.responsibleArea.To) == -1 {
		return true
	}
	return false
}

func (s *SelfConfiguringMerging) Merge(event *models.Event) Merges {
	id := event.ID
	connections := make([][16]byte, 0)
	connections = append(connections, id)
	for _, relArr := range event.Relationships {
		for _, relationship := range relArr {
			connections = append(connections, relationship.ID)
		}
	}

	merges := s.linkEverything(connections)

	log.Debug().Any("event", event).Msg("Added event to knownidentifiers")
	return merges

}

func (s *SelfConfiguringMerging) linkEverything(items [][16]byte) (merges Merges) {

	for _, item := range items {
		if s.responsibleAreaGuard(item) {
			v := s.getOrAddToKnownIdentifiers(item)
			for _, otherItem := range items {
				if item != otherItem {
					v.Linked = append(v.Linked, otherItem)
					//todo: add merge information somewhere

					if s.responsibleAreaGuard(items[0]) { // first must be always event-id
						merges = s.getOrAddMerges(item, otherItem, merges)
					}

				}
			}
		}
	}
	return merges
}

// Todo: adjustment for merging -> only notify of a merge when the event-id is within the responsible area
func (s *SelfConfiguringMerging) getOrAddMerges(to [16]byte, add [16]byte, merges Merges) Merges {
	found := false
	for i := 0; i < len(merges); i++ {
		_, ok := merges[i][to]
		if ok {
			_, ok = merges[i][add]
			if !ok {
				merges[i][add] = true
			}
			found = true
			break
		}
		_, ok = merges[i][add]
		if ok {
			_, ok = merges[i][to]
			if !ok {
				merges[i][to] = true
			}
			found = true
			break
		}
	}
	if !found {
		merges = append(merges, make(map[[16]byte]bool))
		merges[len(merges)-1][to] = true
		merges[len(merges)-1][add] = true
	}
	return merges
}

// Returns the IdentifierReference and a boolean indicating if it was newly created
func (s *SelfConfiguringMerging) getOrAddToKnownIdentifiers(id uuid.UUID) *models.IdentifierReference {
	v, ok := s.knownIdentifiers[id]
	if !ok {
		v = &models.IdentifierReference{
			Self:   id,
			Linked: make([][16]byte, 0),
		}
		s.knownIdentifiers[id] = v
		return v
	}
	return v
}
