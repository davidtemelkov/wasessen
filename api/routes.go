package api

import (
	"net/http"

	"github.com/davidtemelkov/wasessen/assets"
	"github.com/davidtemelkov/wasessen/pages"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes() *chi.Mux {
	r := chi.NewRouter()
	assets.Mount(r)

	r.Get("/", handleServeIndex())

	return r
}

func handleServeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages.Index().Render(r.Context(), w)
	}
}
