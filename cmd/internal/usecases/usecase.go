package usecases

import (
	"agent/cmd/internal/models"
	"agent/cmd/internal/repositories"
	"errors"
	"fmt"
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

func (u UseCases) GetAllUsers() []models.User {
	users := u.repos.User.GetAll()
	return users
}

func (u UseCases) AddNewUser(newUser models.UserRequest) (uuid.UUID, error) {
	exists := u.repos.User.EmailExists(newUser.Email)
	fmt.Println("Exists email: ", exists)
	if exists {
		slog.Error("email already exists", "email", newUser.Email)
		return uuid.Nil, errors.New("email already exists")
	}

	user := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(user)
	fmt.Println("User: ", user)
	return user.ID, nil
}

func (u UseCases) UpdateUser(updateUser models.UserRequest, id string) (models.User, error) {
	exists := u.repos.User.EmailExists(updateUser.Email)
	fmt.Println("Exists email: ", exists)
	if exists {
		slog.Error("email already exists", "email", updateUser.Email)
		return models.User{}, errors.New("email already exists")
	}

	resp, success := u.repos.User.Update(updateUser, id)
	if !success {
		slog.Error("User doesnt exist with this:", "ID", id)
		return models.User{}, errors.New("User doesnt exist with this ID")
	}
	fmt.Println("User update: ", resp)

	return resp, nil
}

func (u UseCases) DeleteUser(id string) (bool, error) {
	success := u.repos.User.Delete(id)
	if !success {
		slog.Error("User doesnt exists:", "ID", id)
		return false, errors.New("User doesnt exists with ID")
	}
	fmt.Println("Deleted user: ", success)

	return true, nil
}
