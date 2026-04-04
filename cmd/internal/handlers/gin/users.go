package ginhandlers

import (
	"agent/cmd/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) RegisterUserEndpoints(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", h.getAllUsers)
		userRoutes.POST("", h.addUser)
		userRoutes.PUT("/:id", h.updateUser)
		userRoutes.DELETE("/:id", h.deleteUser)
	}
}

func (h *Handlers) getAllUsers(c *gin.Context) {
	users := h.usecases.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (h *Handlers) addUser(c *gin.Context) {
	var req models.UserRequest

	// ShouldBindJSON substitui a necessidade do json.NewDecoder e já valida a sintaxe
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := h.usecases.AddNewUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Reason: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.UserResponse{ID: id})
}

func (h *Handlers) updateUser(c *gin.Context) {
	// O router em árvore do Gin garante que a rota PUT /users/:id não atinja
	// este handler se o ID estiver vazio. A validação manual id == "" torna-se obsoleta.
	id := c.Param("id")

	var req models.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Reason: err.Error()})
		return
	}

	user, err := h.usecases.UpdateUser(req, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Reason: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handlers) deleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := h.usecases.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Reason: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
