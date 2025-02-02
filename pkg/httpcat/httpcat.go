package httpcat

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed assets/*
var catFS embed.FS

func SendError(w http.ResponseWriter, code int) {
	filePath := fmt.Sprintf("assets/%d.jpg", code)

	data, err := catFS.ReadFile(filePath)
	if err != nil {
		http.Error(w, "HTTP Cat Image not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
