package entity

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Bio *string `json:"bio"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Bio *string `json:"bio"`
	Email string `json:"email"`
}