package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest : starts server and handle html methods on routes
func HandleRequest() {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/recipes", getRecipes).Methods("GET")
	Router.HandleFunc("/recipes/{recipe}", getRecipe).Methods("GET")
	Router.HandleFunc("/recipes", createrecipe).Methods("POST")
	Router.HandleFunc("/recipes/{recipe}", updateRecipe).Methods("PUT")
	Router.HandleFunc("/recipes/{recipe}", deleteRecipe).Methods("DELETE")
	Router.HandleFunc("/recipes", deleteRecipes).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
