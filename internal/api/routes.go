package api

import (
	"github.com/davidtemelkov/wasessen/internal/assets"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes() *chi.Mux {
	r := chi.NewRouter()
	assets.Mount(r)

	r.Get("/", handleServeIndex())
	r.Get("/recipe/{id}", handleServeRecipe())
	r.Post("/recipe", handleAddRecipe())
	r.Patch("/recipe", handleUpdateRecipe())
	r.Delete("/recipe/{id}", handleRemoveRecipe())
	r.Get("/recipe/modal", handleOpenAddRecipeModal())
	r.Post("/recipequeue", handleAddRecipeToQueue())
	r.Delete("/recipequeue/{id}", handleRemoveRecipeFromQueue())

	return r
}
