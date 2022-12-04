package bookhdl

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rikyhidayat21/book-store-api/dto/bookDto"
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

func (bh *HTTPBookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, err := bh.bookService.Get(id)
	if err != nil {
		shared.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		shared.WriteResponse(w, http.StatusOK, book)
	}
}

func (bh *HTTPBookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// declare variable for request body
	var request bookDto.NewBookRequest
	err := json.NewDecoder(r.Body).Decode(&request) //receive the request from the client
	if err != nil {
		shared.WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		book, appError := bh.bookService.Create(request)
		if appError != nil {
			shared.WriteResponse(w, appError.Code, appError.Message)
		} else {
			shared.WriteResponse(w, http.StatusCreated, book)
		}
	}
}
