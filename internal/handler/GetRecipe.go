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
	var p float64
	params := mux.Vars(req)
	result, err := database.DB.GetRecipe(params["recipe"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := req.URL.Query()
	if v.Get("people") != "" {
		p, err = strconv.ParseFloat(v.Get("people"), 64)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for k := range v {
		if _, ok := result.Ingredients[k]; !ok {
			if k != "people" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
	for key, value := range result.Ingredients {
		if v.Get(key) == "0" {
			queries++
		}
		f := 0.0
		if v.Get(key) != "" {
			f, err = strconv.ParseFloat(v.Get(key), 64)
		}
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		arr2 = append(arr2, float64(value.Quantity))
		arr1 = append(arr1, f)
	}

	for index, v := range arr1 {

		if v > 0 {
			factor = float64(index)
			queries++
		} else if v < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
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
	} else if v.Get("people") == "0" || p < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
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
		i.Ingredients[key].Quantity = int(float64(i.Ingredients[key].Quantity) * factor)
		if i.Ingredients[key].Quantity == 0 {
			i.Ingredients[key].Quantity = 1
		}
	}
	return i
}
