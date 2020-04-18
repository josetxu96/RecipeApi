package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type BreadRecipe struct {
	Name  string  `json:"name" validate:"required"`
	Flour float64 `json:"flour" validate:"required,numeric"`
	Water float64 `json:"water" validate:"required,numeric"`
	Salt  float64 `json:"salt" validate:"required,numeric"`
	Yeast float64 `json:"yeast" validate:"required,numeric"`
	Sugar float64 `json:"sugar" validate:"numeric"`
	Milk  float64 `json:"milk" validate:"numeric"`
}

var breads map[string]BreadRecipe

func factorice(a1, a2 []float64, f int, i BreadRecipe) BreadRecipe {
	var factor float64
	if a2[f] > 0 {
		factor = a1[f] / a2[f]
	}
	i.Flour = i.Flour * factor
	i.Water = i.Water * factor
	i.Salt = i.Salt * factor
	i.Yeast = i.Yeast * factor
	i.Sugar = i.Sugar * factor
	i.Milk = i.Milk * factor
	return i
}

func getBreads(w http.ResponseWriter, req *http.Request) {
	var listBread []BreadRecipe
	for _, v := range breads {
		listBread = append(listBread, v)
	}
	json.NewEncoder(w).Encode(listBread)
}

func getBread(w http.ResponseWriter, req *http.Request) {
	queries := 0
	var factor int
	params := mux.Vars(req)
	result := breads[params["bread"]]
	base := result
	v := req.URL.Query()
	flour, _ := strconv.ParseFloat(v.Get("flour"), 64)
	water, _ := strconv.ParseFloat(v.Get("water"), 64)
	salt, _ := strconv.ParseFloat(v.Get("salt"), 64)
	yeast, _ := strconv.ParseFloat(v.Get("yeast"), 64)
	milk, _ := strconv.ParseFloat(v.Get("milk"), 64)
	sugar, _ := strconv.ParseFloat(v.Get("sugar"), 64)
	arr1 := []float64{flour, water, salt, milk, sugar, yeast}
	arr2 := []float64{base.Flour, base.Water, base.Salt, base.Milk, base.Sugar, base.Yeast}
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
		return
	}
	json.NewEncoder(w).Encode(base)
}

func updateBread(w http.ResponseWriter, req *http.Request) {
	v := validator.New()
	var breadRecipe BreadRecipe
	params := mux.Vars(req)
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if breadRecipe.Name != params["bread"] {
		w.WriteHeader(http.StatusConflict)
		return
	}
	breads[breadRecipe.Name] = breadRecipe
	json.NewEncoder(w).Encode(breads[breadRecipe.Name])
	storeData(breads[breadRecipe.Name], "breads.json")
}

func deleteBread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if _, ok := breads[params["bread"]]; ok {
		delete(breads, params["bread"])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(breads)
	replaceData(breads, "breads.json")
}

func createBread(w http.ResponseWriter, req *http.Request) {
	v := validator.New()
	var breadRecipe BreadRecipe
	_ = json.NewDecoder(req.Body).Decode(&breadRecipe)
	err := v.Struct(breadRecipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if breads[breadRecipe.Name] == (BreadRecipe{}) {
		breads[breadRecipe.Name] = breadRecipe
	} else {
		w.WriteHeader(http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(breads[breadRecipe.Name])
	storeData(breads[breadRecipe.Name], "breads.json")
}

func storeData(s BreadRecipe, f string) {
	_, _ = os.Open(f)
	jsonString, _ := json.Marshal(s)
	ioutil.WriteFile(f, jsonString, 0777)
	return
}

func replaceData(s map[string]BreadRecipe, f string) {
	_, _ = os.Create(f)
	jsonString, _ := json.Marshal(s)
	ioutil.WriteFile(f, jsonString, 0777)
	return
}
func readData(f string) map[string]BreadRecipe {
	file, _ := os.Open(f)
	defer file.Close()
	var breadRecipe map[string]BreadRecipe
	_ = json.NewDecoder(file).Decode(&breadRecipe)
	return breadRecipe
}

func handleRequest() {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/breads", getBreads).Methods("GET")
	Router.HandleFunc("/breads/{bread}", getBread).Methods("GET")
	Router.HandleFunc("/breads/{bread}", updateBread).Methods("PUT")
	Router.HandleFunc("/breads", createBread).Methods("POST")
	Router.HandleFunc("/breads/{bread}", deleteBread).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func main() {
	breads = make(map[string]BreadRecipe)
	breads = readData("breads.json")
	handleRequest()
}
