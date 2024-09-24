package handlers

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		BadRequestHandler(w)
		return
	}

	filePath := "./web" + r.URL.Path
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		NotFoundHandler(w)
		return
	}

	if fileInfo.IsDir() {
		NotFoundHandler(w)
		return
	}

	http.ServeFile(w, r, filePath)
}
