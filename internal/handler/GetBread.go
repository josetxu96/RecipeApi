package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/breadrecipe"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBread(w http.ResponseWriter, req *http.Request) {

	queries := 0
	var factor int
	var arr1 []float64
	var arr2 []float64
	params := mux.Vars(req)
	result, err := database.DB.GetBread(params["bread"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := req.URL.Query()
	for key, value := range result.Ingredients {
		f, _ := strconv.ParseFloat(v.Get(key), 64)
		arr2 = append(arr2, float64(value))
		arr1 = append(arr1, f)
	}

	for index, v := range arr1 {

		if v > 0 {
			factor = index
			queries++
		}

	}

	if queries > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if queries == 1 {
		result = factorice(arr1, arr2, factor, result)
	}
	json.NewEncoder(w).Encode(result)
}

func factorice(a1, a2 []float64, f int, i model.BreadRecipe) model.BreadRecipe {

	var factor float64

	if a2[f] > 0 {
		factor = a1[f] / a2[f]
	}

	for key := range i.Ingredients {
		i.Ingredients[key] = int(float64(i.Ingredients[key]) * factor)
	}
	return i
}
