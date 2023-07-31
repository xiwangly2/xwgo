package main

import (
	"fmt"
	"net/http"

	"xwgo/internal/routes"
)

func main() {
	r := routes.NewRouter()

	fmt.Println("Server started. Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
