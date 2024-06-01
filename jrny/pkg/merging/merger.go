package merging

import (
	"github.com/L4B0MB4/JRNY/jrny/pkg/models"
	"github.com/L4B0MB4/JRNY/jrny/pkg/space"
)

type Merger interface {
	Initialize(*space.ResponsibleArea)
	Merge(*models.Event)
}
