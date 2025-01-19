package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/components"
	"github.com/davidtemelkov/wasessen/internal/data"
	"github.com/davidtemelkov/wasessen/internal/pages"

	"github.com/go-chi/chi/v5"

	"github.com/davidtemelkov/wasessen/internal/utils"
)

func handleAddRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: better error handling, return bad request if malformed input
		newRecipe, err := parseRecipeFromRequest(r)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		err = data.InsertRecipe(r.Context(), newRecipe)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipes
		fmt.Fprintf(w, "recipe added successfully")
	}
}

func handleOpenAddRecipeModal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		components.AddRecipe().Render(r.Context(), w)
	}
}

func handleRemoveRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		err := data.RemoveRecipe(r.Context(), id)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipes
		fmt.Fprintf(w, "recipe removed successfully")
	}
}

func handleUpdateRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: better error handling, return bad request if malformed input
		newRecipe, err := parseRecipeFromRequest(r)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		id := r.FormValue("id")
		oldRecipe, err := data.GetRecipeByID(r.Context(), id)
		if err != nil {
			// TODO: Errors is
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		err = data.UpdateRecipe(r.Context(), oldRecipe, newRecipe)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this rerender recipe details
		fmt.Fprintf(w, "recipe updated successfully")
	}
}

func handleServeRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		recipe, err := data.GetRecipeByID(r.Context(), id)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		pages.Recipe(recipe).Render(r.Context(), w)
	}
}

// TODO: Add form validation
func parseRecipeFromRequest(r *http.Request) (data.Recipe, error) {
	name := r.FormValue("name")
	ingredients := r.FormValue("ingredients")
	preparation := r.FormValue("preparation")
	difficulty := r.FormValue("difficulty")

	file, _, err := r.FormFile("image")
	if err != nil {
		return data.Recipe{}, errors.New("retrieving file error")
	}
	defer file.Close()

	imageURL, err := utils.UploadFile(r.Context(), file)
	if err != nil {
		return data.Recipe{}, errors.New("upload file error")
	}

	newRecipe := data.Recipe{
		Name:        name,
		Ingredients: ingredients,
		Preparation: preparation,
		Difficulty:  difficulty,
		ImageURL:    imageURL,
	}

	return newRecipe, nil
}
