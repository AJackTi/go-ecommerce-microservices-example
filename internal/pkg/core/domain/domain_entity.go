package domain

import (
	uuid "github.com/satori/go.uuid"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/core/events"
)

type IDomainEvent interface {
	events.IEvent
	GetAggregateId() uuid.UUID
	GetAggregateSequenceNumber() int64
	WithAggregate(aggregate uuid.UUID, aggregateSequenceNumber int64)
}

type DomainEvent struct {
	*events.Event
	AggregateId             uuid.UUID `json:"aggregate_id"`
	AggregateSequenceNumber int64     `json:"aggregate_sequence_number"`
}
