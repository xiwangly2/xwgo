package routes

import (
	"net/http"

	"xwgo/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/ua", handlers.UAHandler)

	return mux
}
