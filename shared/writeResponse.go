package shared

import (
	"encoding/json"
	"github.com/rikyhidayat21/book-store-api/logger"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Failed to write response")
		panic(err)
	}
}
