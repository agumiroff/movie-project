package repositories

import (
	"movie-project/internal/app/domain"
	"movie-project/internal/app/repository/models"
)

func userToDomain(user models.User) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Admin:    user.Admin,
	})
}

func domainToUser(user domain.User) models.User {
	return models.User{
		ID:       user.UserID(),
		Username: user.Username(),
		Password: user.Password(),
		Admin:    user.Admin(),
	}
}
