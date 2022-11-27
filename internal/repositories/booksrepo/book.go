package booksrepo

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/rikyhidayat21/book-store-api/exception"
	"github.com/rikyhidayat21/book-store-api/internal/core/domain"
	"github.com/rikyhidayat21/book-store-api/logger"
)

type BookRepository struct {
	client *sqlx.DB
}

// NewBookRepository -> helper to be called when wiring
func NewBookRepository(dbClient *sqlx.DB) *BookRepository {
	return &BookRepository{client: dbClient}
}

// FindAll -> implement interface FindAll()
func (d BookRepository) FindAll() ([]domain.Book, *exception.AppError) {
	// create variable to store books
	books := make([]domain.Book, 0)

	// create error variable
	var err error

	// define query
	findAllSql := "select id, title, year_published, isbn, price, out_of_print, views from books"
	err = d.client.Select(&books, findAllSql)
	if err != nil {
		logger.Error("Error while querying books table " + err.Error())
		return nil, exception.NewUnexpectedError("Unexpected database error")
	}

	return books, nil
}
