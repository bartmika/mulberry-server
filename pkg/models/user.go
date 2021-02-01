// github.com/bartmika/mulberry-server/internal/models/user.go
package models

import (
	"context"
)

// The definition of the user record we will saving in our database.
type User struct {
	Uuid string          `json:"uuid"`
	Name string          `json:"name"`
	Email string         `json:"email"`
	PasswordHash string  `json:"password_hash"`
}

// The interface that *must* be implemented.
type UserRepository interface {
	Create(ctx context.Context, uuid string, name string, email string, passwordHash string) error
	FindByUuid(ctx context.Context, uuid string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

// The struct used to represent the user's `register` POST request data.
type RegisterRequest struct {
	Name string     `json:"name"`
	Email string    `json:"email"`
	Password string `json:"password"`
}

// The struct used to represent the system's response when the `register` POST request was a success.
type RegisterResponse struct {
	Message string `json:"message"`
}

// The struct used to represent the user's `login` POST request data.
type LoginRequest struct {
	Email string    `json:"email"`
	Password string `json:"password"`
}

// The struct used to represent the system's response when the `login` POST request was a success.
type LoginResponse struct {
	AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

// The struct used to represent the user's `refresh token` POST request data.
type RefreshTokenRequest struct {
	Value string     `json:"value"`
}

// The struct used to represent the system's response when the `refresh token` POST request was a success.
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}
