package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getRecipe(w http.ResponseWriter, req *http.Request) {

	queries := 0
	var factor float64
	var arr1 []float64
	var arr2 []float64
	params := mux.Vars(req)
	result, err := database.DB.GetRecipe(params["recipe"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := req.URL.Query()
	p, _ := strconv.ParseFloat(v.Get("people"), 64)

	for key, value := range result.Ingredients {
		f, _ := strconv.ParseFloat(v.Get(key), 64)
		arr2 = append(arr2, float64(value))
		arr1 = append(arr1, f)
	}

	for index, v := range arr1 {

		if v > 0 {
			factor = float64(index)
			queries++
		}

	}

	if queries > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if queries == 1 {
		result = factorice(arr1, arr2, factor, result, false)
		json.NewEncoder(w).Encode(result)
		return
	}
	if p > 0 {
		factor = p / float64(result.People)
		result = factorice(arr1, arr2, factor, result, true)
		result.People = int(p)
	}
	json.NewEncoder(w).Encode(result)
}

func factorice(a1, a2 []float64, f float64, i model.Recipe, people bool) model.Recipe {

	var factor float64
	if people {
		factor = float64(f)
	} else if a2[int(f)] > 0 {
		factor = a1[int(f)] / a2[int(f)]
		i.People = int(float64(i.People) * factor)
	}

	for key := range i.Ingredients {
		i.Ingredients[key] = int(float64(i.Ingredients[key]) * factor)
	}
	return i
}
