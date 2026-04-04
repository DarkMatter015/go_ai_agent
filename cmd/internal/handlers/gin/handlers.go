package ginhandlers

import (
	"agent/cmd/internal/usecases"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	usecases *usecases.UseCases
}

func New(usecases *usecases.UseCases) *Handlers {
	return &Handlers{
		usecases: usecases,
	}
}

func (h *Handlers) Listen(port int) error {
	// Instancia o motor do framework de forma isolada, já com middlewares nativos
	router := gin.Default()

	// Passa a instância explícita para o mapeador de rotas construído no passo anterior
	h.RegisterUserEndpoints(router)

	slog.Info("Listening on", "port", port, "framework", "gin")

	// Inicia o servidor encapsulando a complexidade do http.Server nativo
	return router.Run(fmt.Sprintf(":%v", port))
}
