package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Entity struct {
	id         uuid.UUID
	entityType string
	createdAt  time.Time
	updatedAt  time.Time
}
