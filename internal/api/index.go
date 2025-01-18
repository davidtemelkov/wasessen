package api

import (
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/data"
	"github.com/davidtemelkov/wasessen/internal/pages"
)

func handleServeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := data.GetRecipes(r.Context())
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		recipeQueue, err := data.GetRecipeQueueItems(r.Context())
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		pages.Index(recipes, recipeQueue).Render(r.Context(), w)
	}
}
