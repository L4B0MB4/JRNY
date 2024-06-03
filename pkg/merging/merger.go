package merging

import (
	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/L4B0MB4/JRNY/pkg/space"
)

type Merges []map[[16]byte]bool

type Merger interface {
	Initialize(*space.ResponsibleArea)
	Merge(*models.Event) Merges
}
