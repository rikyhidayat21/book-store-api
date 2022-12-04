package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/rikyhidayat21/book-store-api/internal/core/services/booksvc"
	"github.com/rikyhidayat21/book-store-api/internal/handlers/bookhdl"
	"github.com/rikyhidayat21/book-store-api/internal/repositories/booksrepo"
	"log"
	"net/http"
	"os"
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
	apiV1.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods(http.MethodGet)
	apiV1.HandleFunc("/books", bookHandler.CreateBook).Methods(http.MethodPost)
	apiV1.HandleFunc("/books/{id:[0-9]+}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	// define the connection
	client, err := sqlx.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func dsn() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	dsnSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return dsnSource
}
