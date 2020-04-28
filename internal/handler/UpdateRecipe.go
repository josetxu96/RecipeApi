package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func updateRecipe(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	v := validator.New()
	var recipe model.Recipe
	_ = json.NewDecoder(req.Body).Decode(&recipe)
	err := v.Struct(recipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.DB.UpdateRecipe(&recipe, params["recipe"])

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return
	}

	json.NewEncoder(w).Encode(recipe)
}
