package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/davidtemelkov/wasessen/internal/data"

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
