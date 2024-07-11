package repositories

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"movie-project/internal/app/domain"
	"movie-project/internal/app/repository"
	"movie-project/internal/app/repository/models"
	"time"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	dbUser := domainToUser(user)

	var insertedUser models.User

	err := r.db.QueryRow(ctx, "INSERT INTO users (name, password, admin, created_at, updated_at) values ($1, $2, $3, $4, $5) RETURNING id", dbUser.Username, dbUser.Password, dbUser.Admin, time.Now(), time.Now()).Scan(&insertedUser)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, repository.ErrCodeDuplicateEntry
			}
		}
		return nil, err
	}

	domainUser, err := userToDomain(insertedUser)

	return &domainUser, nil
}
