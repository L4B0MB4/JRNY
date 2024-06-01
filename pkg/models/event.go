package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Type          string                `json:"type" binding:"required"`
	ID            uuid.UUID             `json:"id" binding:"required"`
	Attributes    map[string]string     `json:"attributes"`
	Relationships map[string][]Relation `json:"relationships"`
	ReceivedAt    time.Time
}

type Relation struct {
	Type string    `json:"type" binding:"required"`
	ID   uuid.UUID `json:"id" binding:"required"`
}
