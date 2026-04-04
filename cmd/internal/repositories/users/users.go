package users

import (
	"agent/cmd/internal/models"
	"fmt"
	"sync"
)

type Users struct {
	users []models.User
	mu    sync.RWMutex
}

func New() *Users {
	return &Users{
		users: []models.User{},
	}
}

func (u *Users) GetAll() []models.User {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.users
}

func (u *Users) EmailExists(email string) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, v := range u.users {
		if v.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser models.User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users = append(u.users, newUser)
	fmt.Println("Append user: ", newUser)
	fmt.Println("Users after append: ", u.users)
}

func (u *Users) Update(updateUser models.UserRequest, id string) (models.User, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for i := range u.users {
		if u.users[i].ID.String() == id {
			u.users[i].Name = updateUser.Name
			u.users[i].Email = updateUser.Email
			return u.users[i], true
		}
	}
	return models.User{}, false
}

func (u *Users) Delete(id string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	for i := range u.users {
		if u.users[i].ID.String() == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return true
		}
	}
	return false
}
