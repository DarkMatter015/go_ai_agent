package main

import (
	ginhandlers "agent/cmd/internal/handlers/gin"
	"agent/cmd/internal/repositories"
	"agent/cmd/internal/usecases"
	"log/slog"
	"os"
)

func main() {
	repo := repositories.New()
	useCases := usecases.New(repo)

	// 3. Injeta regra de negócio no framework Gin
	apiHandlers := ginhandlers.New(useCases)

	if err := apiHandlers.Listen(8081); err != nil {
		slog.Error("Falha fatal ao iniciar servidor Gin", "erro", err)
		os.Exit(1)
	}
}
