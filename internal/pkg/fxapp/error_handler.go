package fxapp

import "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"

type FxErrorHandler struct {
	logger logger.Logger
}

func NewFxErrorHandler(logger logger.Logger) *FxErrorHandler {
	return &FxErrorHandler{logger: logger}
}

func (h *FxErrorHandler) HandleError(e error) {
	h.logger.Error(e)
}
