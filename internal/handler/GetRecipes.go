package handler

import (
	"RecipeApi/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func getRecipes(w http.ResponseWriter, req *http.Request) {

	recipe, err := database.DB.GetRecipes()

	if err != nil {

		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(recipe)

}
