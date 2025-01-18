package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/data"
	"github.com/go-chi/chi/v5"
)

func handleAddRecipeToQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: get these form body, add validation
		recipeName := r.FormValue("recipe_name")
		cook := r.FormValue("cook")
		//

		err := data.InsertRecipeQueueItem(r.Context(), recipeName, cook)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipe queue
		fmt.Fprintf(w, "recipe added to queue successfully")
	}
}

func handleRemoveRecipeFromQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		err := data.RemoveRecipeQueueItem(r.Context(), id)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipe queue
		fmt.Fprintf(w, "recipe removed from queue successfully")
	}
}
