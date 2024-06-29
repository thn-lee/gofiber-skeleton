package books

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/gofiber-skelton/pkg/domain"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

type bookUseCase struct {
	bookRepository domain.BookRepository
}

func NewBookUseCase(r domain.BookRepository) domain.BookUseCase {
	return &bookUseCase{
		bookRepository: r,
	}
}

func (u *bookUseCase) DbMigrator() (err error) {
	err = u.bookRepository.DbMigrator()
	return
}

func (u *bookUseCase) CreateBook(book *models.Book) (err error) {
	if len(book.Title) == 0 {
		err = helpers.NewError(fiber.StatusBadRequest, helpers.WhereAmI(), helpers.WhereAmI(), "need: title")
		return
	}
	if len(book.Author) == 0 {
		err = helpers.NewError(fiber.StatusBadRequest, helpers.WhereAmI(), helpers.WhereAmI(), "need: author")
		return
	}
	return u.bookRepository.CreateBook(book)
}

func (u *bookUseCase) EditBook(bookID string, book models.Book) (err error) {
	return u.bookRepository.EditBook(bookID, book)
}

func (u *bookUseCase) DeleteBook(bookID string) (err error) {
	return u.bookRepository.DeleteBook(bookID)
}

func (u *bookUseCase) GetBook(bookID string) (book models.Book, err error) {
	return u.bookRepository.GetBook(bookID)
}

func (u *bookUseCase) GetBooks(criteria models.Book) (books []models.Book, err error) {
	return u.bookRepository.GetBooks(criteria)
}
