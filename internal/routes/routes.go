package routes

import (
	"net/http"

	"xwgo/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", handlers.RootHandler) // Handle root URL
	mux.HandleFunc("/ua", handlers.UAHandler)
	mux.HandleFunc("/qq", handlers.QQHandler)
	mux.HandleFunc("/yiyan", handlers.YiYanDataHandler)
	mux.HandleFunc("/ping", handlers.PingHandler)

	return mux
}
