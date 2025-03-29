package app

import (
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/fxapp/contracts"
)

type OrdersApplicationBuilder struct {
	contracts.ApplicationBuilder
}

// func NewOrdersApplicationBuilder() *OrdersApplicationBuilder {
// 	return &OrdersApplicationBuilder{
// 		fxapp.NewApplicationBuilder(),
// 	}
// }

// func (a *OrdersApplicationBuilder) Build() *OrdersApplicationBuilder {
// 	return NewOrdersApplicationBuilder(
// 		a.GetProviders(),
// 		a.GetDecorates(),
// 		a.Options(),
// 		a.Logger(),
// 		a.Environment(),
// 	)
// }
