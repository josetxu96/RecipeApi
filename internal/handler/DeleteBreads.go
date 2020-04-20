package handler

import (
	"RecipeApi/internal/database"
	"fmt"
	"net/http"
)

func deleteBreads(w http.ResponseWriter, req *http.Request) {
	err := database.DB.DeleteBreads()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
