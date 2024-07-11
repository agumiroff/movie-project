package services

import (
	"context"
	"movie-project/internal/app/domain"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{}
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return s.repo.CreateUser(ctx, user)
}
