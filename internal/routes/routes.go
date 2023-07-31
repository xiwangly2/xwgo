package routes

import (
	"net/http"
	"xwgo/internal/handlers"
)

// Wrapper function to handle panics and return Internal Server Error
func panicHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		handler(w, r)
	}
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes with the panicHandler
	mux.HandleFunc("/", panicHandler(handlers.RootHandler))
	mux.HandleFunc("/ua", panicHandler(handlers.UAHandler))
	mux.HandleFunc("/qq", panicHandler(handlers.QQHandler))
	mux.HandleFunc("/yiyan", panicHandler(handlers.YiYanDataHandler))
	mux.HandleFunc("/ping", panicHandler(handlers.PingHandler))

	return mux
}
