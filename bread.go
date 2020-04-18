package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var DB Store

type BreadRecipe struct {
	Name  string `json:"name" validate:"required"`
	Flour int    `json:"flour" validate:"required,numeric"`
	Water int    `json:"water" validate:"required,numeric"`
	Salt  int    `json:"salt" validate:"required,numeric"`
	Yeast int    `json:"yeast" validate:"required,numeric"`
	Sugar int    `json:"sugar" validate:"numeric"`
	Milk  int    `json:"milk" validate:"numeric"`
}

func factorice(a1, a2 []float64, f int, i BreadRecipe) BreadRecipe {

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

func getBreads(w http.ResponseWriter, req *http.Request) {

	bread, err := DB.GetBreads()

	if err != nil {

		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(bread)

}

func getBread(w http.ResponseWriter, req *http.Request) {

	queries := 0
	var factor int
	params := mux.Vars(req)
	result, err := DB.GetBread(params["bread"])
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

	if base == (BreadRecipe{}) {
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

func createBread(w http.ResponseWriter, req *http.Request) {

	v := validator.New()
	var breadRecipe BreadRecipe
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DB.CreateBread(&breadRecipe)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}

	json.NewEncoder(w).Encode(breadRecipe)
}

func handleRequest() {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/breads", getBreads).Methods("GET")
	Router.HandleFunc("/breads/{bread}", getBread).Methods("GET")
	Router.HandleFunc("/breads", createBread).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func main() {
	connString := "user=jose password=42771618210 dbname=breads sslmode=disable"
	connection, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = connection.Ping()

	if err != nil {
		panic(err)
	}

	DB := database.InitStore(connection)
	handleRequest()

}
