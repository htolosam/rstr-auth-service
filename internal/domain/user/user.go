package user

import (
	"errors"
	"time"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password must have at least 6 characters")
	ErrEmailTaken   = errors.New("email already registered")
)

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}
