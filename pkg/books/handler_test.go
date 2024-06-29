package books_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/gofiber-skelton/mocks"
	"github.com/zercle/gofiber-skelton/pkg/books"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

func TestGetBookHandler(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockResponse := helpers.ResponseForm{
		Success: true,
		Result: models.BooksResponse{
			Books: []models.Book{mockBook},
		},
	}

	mockRepo := new(mocks.BookRepository)

	mockRepo.On("GetBook", mockBook.ID).Return(mockBook, nil)

	app := fiber.New()
	req, err := http.NewRequest("GET", "/api/v1/book/"+mockBook.ID, nil)
	assert.NoError(t, err)

	mockUC := books.NewBookUseCase(mockRepo)
	books.NewBookHandler(app.Group("/api/v1/book"), mockUC)

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockJson, err := json.Marshal(mockResponse)
	assert.NoError(t, err)

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, resp.Body)

	assert.JSONEq(t, string(mockJson), bodyBytes.String())

	mockRepo.AssertExpectations(t)
}

func TestGetBooksHandler(t *testing.T) {
	var mockBookOne, mockBookTwo models.Book
	gofakeit.Struct(&mockBookOne)
	gofakeit.Struct(&mockBookTwo)

	mockResponse := helpers.ResponseForm{
		Success: true,
		Result: models.BooksResponse{
			Books: []models.Book{mockBookOne, mockBookTwo},
		},
	}

	mockRepo := new(mocks.BookRepository)

	mockRepo.On("GetBooks", models.Book{}).Return([]models.Book{mockBookOne, mockBookTwo}, nil)

	app := fiber.New()
	req, err := http.NewRequest("GET", "/api/v1/book/", nil)
	assert.NoError(t, err)

	mockUC := books.NewBookUseCase(mockRepo)
	books.NewBookHandler(app.Group("/api/v1/book"), mockUC)

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockJson, err := json.Marshal(mockResponse)
	assert.NoError(t, err)

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, resp.Body)

	assert.JSONEq(t, string(mockJson), bodyBytes.String())

	mockRepo.AssertExpectations(t)
}
