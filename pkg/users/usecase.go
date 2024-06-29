package users

import (
	"github.com/zercle/gofiber-skelton/pkg/domain"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) GetUser(userID string) (user models.User, err error) {
	return u.userRepo.GetUser(userID)
}

func (u *userUseCase) GetUsers(criteria models.User) (users []models.User, err error) {
	return u.userRepo.GetUsers(criteria)
}

func (u *userUseCase) CreateUser(user *models.User) (err error) {
	return u.userRepo.CreateUser(user)
}

func (u *userUseCase) EditUser(userID string, user models.User) (err error) {
	return u.userRepo.EditUser(userID, user)
}

func (u *userUseCase) DeleteUser(userID string) (err error) {
	return u.userRepo.DeleteUser(userID)
}
