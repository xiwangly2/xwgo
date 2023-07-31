package main

import (
	"fmt"
	"net/http"

	"xwgo/internal/routes"
)

func main() {
	r := routes.NewRouter()

	fmt.Println("Server started. Listening on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
