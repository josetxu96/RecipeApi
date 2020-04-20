package handler

import (
	"RecipeApi/internal/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteBread : calls database to delete a bread
func DeleteBread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := database.DB.DeleteBread(params["bread"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
