package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/breadrecipe"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

// CreateBread : calls database to create a bread and returns it in json
func CreateBread(w http.ResponseWriter, req *http.Request) {

	v := validator.New()
	var breadRecipe model.BreadRecipe
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.DB.CreateBread(&breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return
	}

	json.NewEncoder(w).Encode(breadRecipe)
}
