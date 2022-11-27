package bookhdl

import (
	"github.com/rikyhidayat21/book-store-api/internal/core/ports"
	"github.com/rikyhidayat21/book-store-api/shared"
	"net/http"
)

// HTTPBookHandler inject dependency of ports.BookService
type HTTPBookHandler struct {
	// this will have a dependency of the attribute below
	bookService ports.BookService
}

// NewHTTPBookHandler helper for wiring in app.go
func NewHTTPBookHandler(bookService ports.BookService) *HTTPBookHandler {
	return &HTTPBookHandler{bookService: bookService}
}

// attach the handler to HTTPBookHandler by giving it in the receiver
func (bh *HTTPBookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAll()
	if err != nil {
		shared.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		shared.WriteResponse(w, http.StatusOK, books)
	}
}
