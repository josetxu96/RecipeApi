package handler

import (
	"RecipeApi/internal/database"
	model "RecipeApi/internal/model/breadRecipe"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type Handler interface {
	GetBreads(w http.ResponseWriter, req *http.Request)
	GetBread(w http.ResponseWriter, req *http.Request)
	CreateBread(w http.ResponseWriter, req *http.Request)
}

type store struct {
	db database.Store
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

func (db *store) GetBreads(w http.ResponseWriter, req *http.Request) {

	bread, err := db.db.GetBreads()

	if err != nil {

		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(bread)

}

func (db *store) GetBread(w http.ResponseWriter, req *http.Request) {

	queries := 0
	var factor int
	params := mux.Vars(req)
	result, err := db.db.GetBread(params["bread"])
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

func (db *store) CreateBread(w http.ResponseWriter, req *http.Request) {

	v := validator.New()
	var breadRecipe model.BreadRecipe
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.db.CreateBread(&breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}

	json.NewEncoder(w).Encode(breadRecipe)
}

func (db *store) DeleteBread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := db.db.DeleteBread(params["bread"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func InitHandler(_store database.Store) Handler {
	return &store{db: _store}
}
