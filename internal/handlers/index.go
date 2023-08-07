package handlers

import (
	"net/http"
	"os"
)

var indexHTML []byte

func init() {
	loadIndexHTML()
}

func loadIndexHTML() {
	file, err := os.ReadFile("static/index.html")
	if err != nil {
		panic("Error reading index.html: " + err.Error())
	}
	indexHTML = file
}

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write(indexHTML)
	if err != nil {
		return
	}
}
