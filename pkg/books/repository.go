package books

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/gofiber-skelton/pkg/constants"
	"github.com/zercle/gofiber-skelton/pkg/domain"
	"github.com/zercle/gofiber-skelton/pkg/models"
	"gorm.io/gorm"
)

type bookRepository struct {
	mainDbConn *gorm.DB
}

func NewBookRepository(mainDbConn *gorm.DB) domain.BookRepository {
	return &bookRepository{
		mainDbConn: mainDbConn,
	}
}

func (r *bookRepository) DbMigrator() (err error) {
	err = r.mainDbConn.AutoMigrate(&models.Book{})
	return
}

func (r *bookRepository) CreateBook(book *models.Book) (err error) {
	if r.mainDbConn == nil {
		err = helpers.NewError(fiber.StatusServiceUnavailable, helpers.WhereAmI(), constants.ErrDBGone)
		return
	}

	dbTx := r.mainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&models.Book{})

	if err = dbTx.Create(book).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error
	return
}

func (r *bookRepository) EditBook(bookID string, book models.Book) (err error) {
	if r.mainDbConn == nil {
		err = helpers.NewError(fiber.StatusServiceUnavailable, helpers.WhereAmI(), constants.ErrDBGone)
		return
	}

	dbTx := r.mainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&models.Book{})
	dbTx = dbTx.Where(models.Book{ID: bookID})

	if err = dbTx.Updates(book).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error

	return
}

func (r *bookRepository) DeleteBook(bookID string) (err error) {
	if r.mainDbConn == nil {
		err = helpers.NewError(fiber.StatusServiceUnavailable, helpers.WhereAmI(), constants.ErrDBGone)
		return
	}

	dbTx := r.mainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&models.Book{})
	dbTx = dbTx.Where(models.Book{ID: bookID})

	if err = dbTx.Delete(&models.Book{}).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error

	return
}

func (r *bookRepository) GetBook(bookID string) (book models.Book, err error) {
	if r.mainDbConn == nil {
		err = helpers.NewError(fiber.StatusServiceUnavailable, helpers.WhereAmI(), constants.ErrDBGone)
		return
	}

	dbTx := r.mainDbConn.Model(&models.Book{})
	dbTx = dbTx.Where(models.Book{ID: bookID})
	err = dbTx.Take(&book).Error

	return
}

func (r *bookRepository) GetBooks(criteria models.Book) (books []models.Book, err error) {
	if r.mainDbConn == nil {
		err = helpers.NewError(fiber.StatusServiceUnavailable, helpers.WhereAmI(), constants.ErrDBGone)
		return
	}

	dbTx := r.mainDbConn.Model(&models.Book{})

	if len(criteria.ID) != 0 {
		dbTx = dbTx.Where(models.Book{ID: criteria.ID})
	} else {
		if len(criteria.Title) != 0 {
			dbTx = dbTx.Where("title LIKE ?", "%"+criteria.Title+"%")
		}

		if len(criteria.Author) != 0 {
			dbTx = dbTx.Where("author LIKE ?", "%"+criteria.Author+"%")
		}
	}

	err = dbTx.Find(&books).Error

	return
}

func (r *bookRepository) ImportBooks(books []models.Book) (errs []error) {

	errCh := make(chan error, (runtime.NumCPU()/2)+1)
	defer close(errCh)

	for _, book := range books {
		go r.importBookChannel(book, errCh)
		if err := <-errCh; err != nil {
			errs = append(errs, err)
		}
	}

	return
}

func (r *bookRepository) importBookChannel(book models.Book, errCh chan error) {
	errCh <- r.CreateBook(&book)
}
