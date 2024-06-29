package books

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/gofiber-skelton/pkg/domain"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

type BookHandler struct {
	bookUseCase domain.BookUseCase
}

func NewBookHandler(bookRoute fiber.Router, bookUseCase domain.BookUseCase) {

	handler := &BookHandler{
		bookUseCase: bookUseCase,
	}

	bookRoute.Get("/:bookID?", handler.GetBooks())
}

func (h *BookHandler) GetBooks() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		bookID := c.Params("bookID")
		if len(bookID) != 0 {
			book, err := h.bookUseCase.GetBook(bookID)
			if err != nil {
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    fiber.StatusServiceUnavailable,
					Message: err.Error(),
				})
			}
			responseForm.Result = models.BooksResponse{
				Books: []models.Book{book},
			}
		} else {
			criteria := models.Book{}
			criteria.Title = c.FormValue("title")
			criteria.Author = c.FormValue("author")
			books, err := h.bookUseCase.GetBooks(criteria)
			if err != nil {
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    fiber.StatusServiceUnavailable,
					Message: err.Error(),
				})
			}
			responseForm.Result = models.BooksResponse{
				Books: books,
			}
		}

		if len(responseForm.Errors) == 0 {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
