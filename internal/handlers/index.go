package handlers

import (
	"io/ioutil"
	"net/http"
)

var indexHTML []byte

func init() {
	loadIndexHTML()
}

func loadIndexHTML() {
	file, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic("Error reading index.html: " + err.Error())
	}
	indexHTML = file
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(indexHTML)
}
