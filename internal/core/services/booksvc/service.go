package booksvc

import (
	"github.com/rikyhidayat21/book-store-api/dto/bookDto"
	"github.com/rikyhidayat21/book-store-api/exception"
	"github.com/rikyhidayat21/book-store-api/internal/core/domain"
	"github.com/rikyhidayat21/book-store-api/internal/core/ports"
)

type service struct {
	repo ports.BookRepository
}

// NewBookService helper to be called when wiring
func NewBookService(repository ports.BookRepository) *service {
	return &service{repo: repository}
}

// GetAll -> implement service
func (s *service) GetAll() ([]domain.Book, *exception.AppError) {
	return s.repo.FindAll()
}

func (s *service) Get(id string) (*bookDto.BookResponse, *exception.AppError) {
	b, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := b.ToDto()

	return &response, nil
}
