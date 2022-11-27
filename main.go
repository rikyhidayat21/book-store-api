package main

import (
	"github.com/rikyhidayat21/book-store-api/app"
	"github.com/rikyhidayat21/book-store-api/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
