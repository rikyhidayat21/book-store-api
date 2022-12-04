// Package ports - Primary Ports
package ports

import (
	"github.com/rikyhidayat21/book-store-api/dto/bookDto"
	"github.com/rikyhidayat21/book-store-api/exception"
)

type BookService interface {
	GetAll() ([]bookDto.BookResponse, *exception.AppError)
	Get(string) (*bookDto.BookResponse, *exception.AppError)
	Create(request bookDto.NewBookRequest) (*bookDto.NewBookResponse, *exception.AppError)
}
