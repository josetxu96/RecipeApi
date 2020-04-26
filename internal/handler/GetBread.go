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
	params := mux.Vars(req)
	result, err := database.DB.GetBread(params["bread"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	base := result
	v := req.URL.Query()
	flour, _ := strconv.ParseFloat(v.Get("flour"), 64)
	water, _ := strconv.ParseFloat(v.Get("water"), 64)
	salt, _ := strconv.ParseFloat(v.Get("salt"), 64)
	yeast, _ := strconv.ParseFloat(v.Get("yeast"), 64)
	milk, _ := strconv.ParseFloat(v.Get("milk"), 64)
	sugar, _ := strconv.ParseFloat(v.Get("sugar"), 64)
	arr1 := []float64{flour, water, salt, milk, sugar, yeast}
	arr2 := []float64{float64(base.Ingredients.Flour), float64(base.Ingredients.Water), float64(base.Ingredients.Salt), float64(base.Ingredients.Milk), float64(base.Ingredients.Sugar), float64(base.Ingredients.Yeast)}

	if base == (model.BreadRecipe{}) {
		w.WriteHeader(http.StatusNotFound)
		return
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
		result = factorice(arr1, arr2, factor, base)
		json.NewEncoder(w).Encode(result)
	}
	json.NewEncoder(w).Encode(base)
}

func factorice(a1, a2 []float64, f int, i model.BreadRecipe) model.BreadRecipe {

	var factor float64

	if a2[f] > 0 {
		factor = a1[f] / a2[f]
	}

	i.Ingredients.Flour = int(float64(i.Ingredients.Flour) * factor)
	i.Ingredients.Water = int(float64(i.Ingredients.Water) * factor)
	i.Ingredients.Salt = int(float64(i.Ingredients.Salt) * factor)
	i.Ingredients.Yeast = int(float64(i.Ingredients.Yeast) * factor)
	i.Ingredients.Sugar = int(float64(i.Ingredients.Sugar) * factor)
	i.Ingredients.Milk = int(float64(i.Ingredients.Milk) * factor)

	return i
}
