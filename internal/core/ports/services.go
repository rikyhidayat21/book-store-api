// Package ports - Primary Ports
package ports

import (
	"github.com/rikyhidayat21/book-store-api/dto/bookDto"
	"github.com/rikyhidayat21/book-store-api/exception"
	"github.com/rikyhidayat21/book-store-api/internal/core/domain"
)

type BookService interface {
	GetAll() ([]domain.Book, *exception.AppError)
	Get(string) (*bookDto.BookResponse, *exception.AppError)
}
