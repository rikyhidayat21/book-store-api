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
func (s *service) GetAll() ([]bookDto.BookResponse, *exception.AppError) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]bookDto.BookResponse, 0)
	for _, b := range books {
		response = append(response, b.ToDto())
	}
	return response, nil
}

func (s *service) Get(id string) (*bookDto.BookResponse, *exception.AppError) {
	b, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := b.ToDto()

	return &response, nil
}

func (s *service) Create(request bookDto.NewBookRequest) (*bookDto.NewBookResponse, *exception.AppError) {
	// validate

	// transform the request to be a DTO
	b := domain.Book{
		Title:         request.Title,
		YearPublished: request.YearPublished,
		Isbn:          request.Isbn,
		Price:         request.Price,
		OutOfPrint:    request.OutOfPrint,
		Views:         request.Views,
	}

	// pass variable
	newBook, err := s.repo.Save(b)
	if err != nil {
		return nil, err
	}

	// transform to new account response
	responseDto := newBook.ToNewBookResponseDto()
	return &responseDto, nil
}

func (s *service) Delete(id string) *exception.AppError {
	_, err := s.repo.ById(id)
	if err != nil {
		return err
	}

	err = s.repo.Destroy(id)
	if err != nil {
		return err
	}

	return nil
}
