package assets

import (
	"embed"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

//go:embed dist/*
var Assets embed.FS

func Mount(r chi.Router) {
	r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		file := chi.URLParam(r, "*")

		w.Header().Set("Content-Type", getMimeType(file))
		http.FileServer(http.FS(Assets)).ServeHTTP(w, r)
	})

}

func getMimeType(filename string) string {
	if strings.HasSuffix(filename, ".css") {
		return "text/css"
	}
	if strings.HasSuffix(filename, ".svg") {
		return "image/svg+xml"
	}
	return "application/octet-stream"
}
