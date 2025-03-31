package store

import "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/es/models"

// AggregateStore is responsible for loading and saving Aggregate.
type AggregateStore[T models.IHaveEventSourcedAggregate] interface {
}
