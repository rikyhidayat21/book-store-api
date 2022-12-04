// Package ports - Secondary Ports
package ports

import (
	"github.com/rikyhidayat21/book-store-api/exception"
	"github.com/rikyhidayat21/book-store-api/internal/core/domain"
)

type BookRepository interface {
	FindAll() ([]domain.Book, *exception.AppError)
	ById(string) (*domain.Book, *exception.AppError) // why using pointer? because we want to send `nil` in case there's no book available
	Save(book domain.Book) (*domain.Book, *exception.AppError)
}
