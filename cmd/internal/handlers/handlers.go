package handlers

import "agent/cmd/internal/usecases"

type Handlers struct {
	usecases *usecases.UseCases
}

func New(usecases *usecases.UseCases) *Handlers {
	return &Handlers{
		usecases: usecases,
	}
}