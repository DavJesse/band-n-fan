package handlers

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		badRequestHandler(w)
		return
	}

	filePath := "." + r.URL.Path
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		notFoundHandler(w)
		return
	}

	if fileInfo.IsDir() {
		notFoundHandler(w)
		return
	}
	http.ServeFile(w, r, filePath)
}
