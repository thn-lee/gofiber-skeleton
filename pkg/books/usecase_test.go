package books_test

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zercle/gofiber-skelton/mocks"
	"github.com/zercle/gofiber-skelton/pkg/books"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

func TestGetBookUseCase(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBookRepo := new(mocks.BookRepository)
	mockBookUseCase := new(mocks.BookUseCase)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetBook", mock.AnythingOfType("string")).Return(mockBook, nil).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		result, err := useCase.GetBook(mockBook.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockBook, result)

		mockBookUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("GetBook", mock.AnythingOfType("string")).Return(models.Book{}, errors.New("call error")).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		result, err := useCase.GetBook(mockBook.ID)

		assert.Error(t, err)
		assert.Equal(t, models.Book{}, result)

		mockBookUseCase.AssertExpectations(t)
	})

}

func TestGetBooksUseCase(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBooks := []models.Book{mockBook}

	mockBookRepo := new(mocks.BookRepository)
	mockBookUseCase := new(mocks.BookUseCase)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetBooks", mock.Anything).Return(mockBooks, nil).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		result, err := useCase.GetBooks(mockBook)

		assert.NoError(t, err)
		assert.Equal(t, mockBooks, result)

		mockBookUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("GetBooks", mock.Anything).Return([]models.Book{}, errors.New("call error")).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		result, err := useCase.GetBooks(mockBook)

		assert.Error(t, err)
		assert.Empty(t, result)

		mockBookUseCase.AssertExpectations(t)
	})

}

func TestCreateBookUseCase(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBookRepo := new(mocks.BookRepository)
	mockBookUseCase := new(mocks.BookUseCase)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("CreateBook", mock.Anything).Return(nil).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.CreateBook(&mockBook)

		assert.NoError(t, err)

		mockBookUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("CreateBook", mock.Anything).Return(errors.New("call error")).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.CreateBook(&mockBook)

		assert.Error(t, err)

		mockBookUseCase.AssertExpectations(t)
	})

}

func TestEditBookUseCase(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBookRepo := new(mocks.BookRepository)
	mockBookUseCase := new(mocks.BookUseCase)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("EditBook", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.EditBook(mockBook.ID, mockBook)

		assert.NoError(t, err)

		mockBookUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("EditBook", mock.AnythingOfType("string"), mock.Anything).Return(errors.New("call error")).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.EditBook(mockBook.ID, mockBook)

		assert.Error(t, err)

		mockBookUseCase.AssertExpectations(t)
	})

}

func TestDeleteBookUseCase(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBookRepo := new(mocks.BookRepository)
	mockBookUseCase := new(mocks.BookUseCase)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("DeleteBook", mock.AnythingOfType("string")).Return(nil).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.DeleteBook(mockBook.ID)

		assert.NoError(t, err)

		mockBookUseCase.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("DeleteBook", mock.AnythingOfType("string")).Return(errors.New("call error")).Once()

		useCase := books.NewBookUseCase(mockBookRepo)

		err := useCase.DeleteBook(mockBook.ID)

		assert.Error(t, err)

		mockBookUseCase.AssertExpectations(t)
	})

}
