package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest : starts server and handle html methods on routes
func HandleRequest() {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/breads", getBreads).Methods("GET")
	Router.HandleFunc("/breads/{bread}", getBread).Methods("GET")
	Router.HandleFunc("/breads", createBread).Methods("POST")
	Router.HandleFunc("/breads/{bread}", updateBread).Methods("PUT")
	Router.HandleFunc("/breads/{bread}", deleteBread).Methods("DELETE")
	Router.HandleFunc("/breads", deleteBreads).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
