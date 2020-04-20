package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/breadrecipe"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func updateBread(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	v := validator.New()
	var breadRecipe model.BreadRecipe
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.DB.UpdateBread(&breadRecipe, params["bread"])

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return
	}

	json.NewEncoder(w).Encode(breadRecipe)
}
