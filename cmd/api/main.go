package main

import (
	"agent/cmd/internal/handlers"
	"agent/cmd/internal/repositories"
	"agent/cmd/internal/usecases"
)

func main() {
	// Injeção de Dependências
	repositories := repositories.New()
	useCases := usecases.New(repositories)
	handlers := handlers.New(useCases)

	handlers.Listen(8080)
}
