package usecases

import (
	"agent/cmd/internal/models"
	"agent/cmd/internal/repositories"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) GetAllUsers() ([]models.User, error) {
	users := u.repos.User.GetAll()
	return users, nil
}

func (u UseCases) AddNewUser(newUser models.UserRequest) (uuid.UUID, error) {
	if u.repos.User.EmailExists(newUser.Email) {
		slog.Error("email already exists", "email", newUser.Email)
		return uuid.Nil, errors.New("email already exists")
	}

	user := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(user)

	return user.ID, nil
}
