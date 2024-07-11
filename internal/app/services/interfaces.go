package services

import (
	"context"
	"movie-project/internal/app/domain"
)

type UserRepository interface {
	//GetUser(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	//GetUserByID(ctx context.Context, id int) (domain.User, error)
}
