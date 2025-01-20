package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/components"
	"github.com/davidtemelkov/wasessen/internal/data"
	"github.com/go-chi/chi/v5"
)

func handleAddRecipeToQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeName := r.FormValue("recipe_name")
		cook := r.FormValue("cook")

		updatedRecipeQueue, err := data.AddRecipeToQueue(r.Context(), recipeName, cook)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		recipeQueueHTML := components.RecipeQueue(updatedRecipeQueue.Queue).Render(r.Context(), w)
		fmt.Fprint(w, recipeQueueHTML)
	}
}

func handleRemoveRecipeFromQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		updatedRecipeQueue, err := data.RemoveRecipeFromQueue(r.Context(), id)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		recipeQueueHTML := components.RecipeQueue(updatedRecipeQueue.Queue).Render(r.Context(), w)
		fmt.Fprint(w, recipeQueueHTML)
	}
}

func handleMoveRecipeInQueue(up bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		updatedRecipeQueue, err := data.MoveRecipeInQueue(r.Context(), id, up)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		recipeQueueHTML := components.RecipeQueue(updatedRecipeQueue.Queue).Render(r.Context(), w)
		fmt.Fprint(w, recipeQueueHTML)
	}
}
