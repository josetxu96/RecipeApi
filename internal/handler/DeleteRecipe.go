package handler

import (
	"RecipeApi/internal/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func deleteRecipe(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := database.DB.DeleteRecipe(params["recipe"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
