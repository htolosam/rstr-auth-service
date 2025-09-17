package service

import (
	"context"
	"github.com/htolosam/rstr-auth-service/internal/crypto"
	"github.com/htolosam/rstr-auth-service/internal/domain/user"
	"strings"
	"time"
)

type IUserRepository interface {
	Save(ctx context.Context, user user.User) error
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

type AuthService struct {
	repo IUserRepository
}

func NewAuthService(repo IUserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, email, password, role string) (*user.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if !strings.Contains(email, "@") {
		return nil, user.ErrInvalidEmail
	}
	if len(password) < 6 {
		return nil, user.ErrWeakPassword
	}

	exists, _ := s.repo.FindByEmail(ctx, email)
	if exists != nil {
		return nil, user.ErrEmailTaken
	}

	hash, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &user.User{
		Email:        email,
		PasswordHash: hash,
		Role:         role,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.Save(ctx, *u); err != nil {
		return nil, err
	}

	return u, nil
}
