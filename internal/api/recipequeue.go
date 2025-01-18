package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/data"
)

func handleAddRecipeToQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeName := r.FormValue("recipe_name")
		cook := r.FormValue("cook")

		err := data.InsertRecipeQueueItem(r.Context(), recipeName, cook)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipe queue
		fmt.Fprintf(w, "recipe added to queue successfully")
	}
}
