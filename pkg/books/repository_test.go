package books_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/zercle/gofiber-skelton/pkg/books"
	"github.com/zercle/gofiber-skelton/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetBookRepo(t *testing.T) {
	var mockBook models.Book
	gofakeit.Struct(&mockBook)

	mockBook.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	mockBook.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	mockBook.DeletedAt = gorm.DeletedAt(sql.NullTime{Valid: false})

	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "description", "author", "title", "created_at", "updated_at", "deleted_at"}).AddRow(mockBook.ID, mockBook.Description, mockBook.Author, mockBook.Title, mockBook.CreatedAt, mockBook.UpdatedAt, mockBook.DeletedAt)

	queryRegexp := "^SELECT (.+) FROM `books` (.+)$"

	mock.ExpectQuery(queryRegexp).WillReturnRows(rows)

	t.Run("success", func(t *testing.T) {
		mockRepo := books.NewBookRepository(gdb)

		result, err := mockRepo.GetBook(mockBook.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, mockBook, result)
	})

	t.Run("fail-db-conn", func(t *testing.T) {
		mockRepo := books.NewBookRepository(nil)

		result, err := mockRepo.GetBook(mockBook.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
