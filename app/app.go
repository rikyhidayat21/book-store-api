package app

import "net/http"

func Start() {
	// define mux

	http.ListenAndServe("localhost:8000", nil)
}
