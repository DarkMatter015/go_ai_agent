package main

import (
	"agent/cmd/internal/handlers/native"
	"agent/cmd/internal/repositories"
	"agent/cmd/internal/usecases"
	"log/slog"
	"os"
)

func main() {
	repositories := repositories.New()
	useCases := usecases.New(repositories)
	apiHandlers := native.New(useCases)

	if err := apiHandlers.Listen(8080); err != nil {
		slog.Error("Falha fatal ao iniciar servidor nativo", "erro", err)
		os.Exit(1)
	}
}
