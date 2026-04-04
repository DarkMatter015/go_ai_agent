package repositories

import (
	"agent/cmd/internal/models"
	"agent/cmd/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(user models.User)
		Update(updateUser models.UserRequest, id string) (models.User, bool)
		Delete(id string) bool
		EmailExists(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}
