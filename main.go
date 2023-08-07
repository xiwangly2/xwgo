package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"xwgo/internal/routes"
)

func isPortAvailable(port string) bool {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)
	return true
}

func main() {
	port := "8080"

	if !isPortAvailable(port) {
		fmt.Printf("Port %s is already in use. Please choose a different port.\n", port)
		return
	}

	r := routes.NewRouter()

	fmt.Printf("Server starting. Listening on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Printf("Port %s is already in use. Please choose a different port.\n", port)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}
