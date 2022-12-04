package booksrepo

import (
	"database/sql"
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

func (d BookRepository) ById(id string) (*domain.Book, *exception.AppError) {
	// create variable with default value of domain book
	var b domain.Book

	// define query
	findByIdSql := "select id, title, year_published, isbn, price, out_of_print, views from books where id = ?"

	// call the sqlx
	err := d.client.Get(&b, findByIdSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("Book not found " + err.Error())
			return nil, exception.NewNotFoundError("Book not found")
		} else {
			logger.Error("Error while scanning book " + err.Error())
			return nil, exception.NewUnexpectedError("Unexpected database error")
		}
	}

	return &b, nil
}

func (d BookRepository) Save(b domain.Book) (*domain.Book, *exception.AppError) {
	sqlInsert := "INSERT INTO books (title, year_published, isbn, price, out_of_print, views) VALUES (?, ?, ?, ?, ?, ?)"

	// execute the query
	result, err := d.client.Exec(sqlInsert, b.Title, b.YearPublished, b.Isbn, b.Price, b.OutOfPrint, b.Views)
	if err != nil {
		logger.Error("Error while creating new book: " + err.Error())
		return nil, exception.NewUnexpectedError("Unexpected error from database")
	}

	// return the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id for new book: " + err.Error())
		return nil, exception.NewUnexpectedError("Unexpected error from database")
	}

	b.Id = id
	return &b, nil
}

func (d BookRepository) Destroy(id string) *exception.AppError {
	// define query
	destroySql := "delete from books where id = ?"

	_, err := d.client.Exec(destroySql, id)
	if err != nil {
		logger.Error("Error while destroying book " + err.Error())
		return exception.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}
