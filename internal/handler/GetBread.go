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
	arr2 := []float64{float64(base.Flour), float64(base.Water), float64(base.Salt), float64(base.Milk), float64(base.Sugar), float64(base.Yeast)}

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

	i.Flour = int(float64(i.Flour) * factor)
	i.Water = int(float64(i.Water) * factor)
	i.Salt = int(float64(i.Salt) * factor)
	i.Yeast = int(float64(i.Yeast) * factor)
	i.Sugar = int(float64(i.Sugar) * factor)
	i.Milk = int(float64(i.Milk) * factor)

	return i
}
