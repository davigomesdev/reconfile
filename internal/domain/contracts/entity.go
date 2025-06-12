package contracts

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (e *Entity) Init(id *string, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) {
	now := time.Now()

	if id != nil && *id != "" {
		e.ID = *id
	} else {
		e.ID = uuid.New().String()
	}

	if createdAt != nil {
		e.CreatedAt = *createdAt
	} else {
		e.CreatedAt = now
	}

	if updatedAt != nil {
		e.UpdatedAt = *updatedAt
	} else {
		e.UpdatedAt = now
	}

	e.DeletedAt = deletedAt
}
