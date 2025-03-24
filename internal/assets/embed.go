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

	// For local development, if needed
	// r.Route("/dist", func(r chi.Router) {
	// 	r.Use(func(next http.Handler) http.Handler {
	// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 			next.ServeHTTP(w, r)
	// 		})
	// 	})

	// 	r.Handle("/*", http.FileServer(http.FS(Assets)))
	// })
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
