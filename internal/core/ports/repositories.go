package ports

import (
	"github.com/rikyhidayat21/book-store-api/exception"
	"github.com/rikyhidayat21/book-store-api/internal/core/domain"
)

type BookRepository interface {
	FindAll() ([]domain.Book, *exception.AppError)
}
