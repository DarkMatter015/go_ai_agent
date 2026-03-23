package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Email string    `json:"email"`
}

type UserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
