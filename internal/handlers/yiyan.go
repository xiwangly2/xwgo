package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	yiYanData []string
	mutex     sync.Mutex
)

func init() {
	filename := "static/yiyandata.dat"

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	yiYanData = strings.Split(string(data), "\n")
	for i, line := range yiYanData {
		yiYanData[i] = strings.TrimSpace(line)
	}
}

func YiYanDataHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	randomLine := yiYanData[rand.Intn(len(yiYanData))]
	mutex.Unlock()

	if r.URL.Query().Get("return") == "json" {
		data := map[string]interface{}{
			"code": "200",
			"text": randomLine,
		}

		dataJSON, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataJSON)
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte(randomLine))
		if err != nil {
			return
		}
	}
}
