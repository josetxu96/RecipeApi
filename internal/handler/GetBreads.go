package handler

import (
	"RecipeApi/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetBreads : calls db to get all breads and returns them in json
func GetBreads(w http.ResponseWriter, req *http.Request) {

	bread, err := database.DB.GetBreads()

	if err != nil {

		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(bread)

}
