package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rikyhidayat21/book-store-api/internal/core/services/booksvc"
	"github.com/rikyhidayat21/book-store-api/internal/handlers/bookhdl"
	"github.com/rikyhidayat21/book-store-api/internal/repositories/booksrepo"
	"log"
	"net/http"
	"time"
)

func Start() {
	// define mux
	router := mux.NewRouter()

	// define db
	dbClient := getDbClient()

	// WIRING BEGIN
	// repositories
	bookRepository := booksrepo.NewBookRepository(dbClient)

	// services
	bookService := booksvc.NewBookService(bookRepository)

	// handlers
	bookHandler := bookhdl.NewHTTPBookHandler(bookService)

	// define routes
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/books", bookHandler.GetAllBooks).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	// define the connection
	client, err := sqlx.Open("mysql", "root:root1234@tcp(localhost:3306)/book-store")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
