package users_test

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zercle/gofiber-skelton/mocks"
	"github.com/zercle/gofiber-skelton/pkg/models"
	"github.com/zercle/gofiber-skelton/pkg/users"
)

func TestGetUserUseCase(t *testing.T) {
	var mockUser models.User
	gofakeit.Struct(&mockUser)

	mockUserRepo := new(mocks.UserRepository)
	mockUserUseCase := new(mocks.UserUseCase)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUser", mock.AnythingOfType("string")).Return(mockUser, nil).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		result, err := useCase.GetUser(mockUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)

		mockUserUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetUser", mock.AnythingOfType("string")).Return(models.User{}, errors.New("call error")).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		result, err := useCase.GetUser(mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, models.User{}, result)

		mockUserUseCase.AssertExpectations(t)
	})

}

func TestGetUsersUseCase(t *testing.T) {
	var mockUser models.User
	gofakeit.Struct(&mockUser)

	mockUsers := []models.User{mockUser}

	mockUserRepo := new(mocks.UserRepository)
	mockUserUseCase := new(mocks.UserUseCase)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUsers", mock.Anything).Return(mockUsers, nil).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		result, err := useCase.GetUsers(mockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUsers, result)

		mockUserUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetUsers", mock.Anything).Return([]models.User{}, errors.New("call error")).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		result, err := useCase.GetUsers(mockUser)

		assert.Error(t, err)
		assert.Empty(t, result)

		mockUserUseCase.AssertExpectations(t)
	})

}

func TestCreateUserUseCase(t *testing.T) {
	var mockUser models.User
	gofakeit.Struct(&mockUser)

	mockUserRepo := new(mocks.UserRepository)
	mockUserUseCase := new(mocks.UserUseCase)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.CreateUser(&mockUser)

		assert.NoError(t, err)

		mockUserUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(errors.New("call error")).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.CreateUser(&mockUser)

		assert.Error(t, err)

		mockUserUseCase.AssertExpectations(t)
	})

}

func TestEditUserUseCase(t *testing.T) {
	var mockUser models.User
	gofakeit.Struct(&mockUser)

	mockUserRepo := new(mocks.UserRepository)
	mockUserUseCase := new(mocks.UserUseCase)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("EditUser", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.EditUser(mockUser.ID, mockUser)

		assert.NoError(t, err)

		mockUserUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("EditUser", mock.AnythingOfType("string"), mock.Anything).Return(errors.New("call error")).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.EditUser(mockUser.ID, mockUser)

		assert.Error(t, err)

		mockUserUseCase.AssertExpectations(t)
	})

}

func TestDeleteUserUseCase(t *testing.T) {
	var mockUser models.User
	gofakeit.Struct(&mockUser)

	mockUserRepo := new(mocks.UserRepository)
	mockUserUseCase := new(mocks.UserUseCase)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("DeleteUser", mock.AnythingOfType("string")).Return(nil).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.DeleteUser(mockUser.ID)

		assert.NoError(t, err)

		mockUserUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("DeleteUser", mock.AnythingOfType("string")).Return(errors.New("call error")).Once()

		useCase := users.NewUserUseCase(mockUserRepo)

		err := useCase.DeleteUser(mockUser.ID)

		assert.Error(t, err)

		mockUserUseCase.AssertExpectations(t)
	})

}
