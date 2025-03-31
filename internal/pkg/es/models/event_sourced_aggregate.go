package models

import "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/core/domain"

// IHaveEventSourcedAggregate this interface should implement by actual aggregate root class in our domain_events
type IHaveEventSourcedAggregate interface {
}

type WhenFunc func(event domain.IDomainEvent) error

// EventSourcedAggregateRoot base aggregate contains all main necessary fields
type EventSourcedAggregateRoot struct {
	*domain.Entity
	originalVersion   int64
	currentVersion    int64
	uncommittedEvents []domain.IDomainEvent
	when              WhenFunc
}
