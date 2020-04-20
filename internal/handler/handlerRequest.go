package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type handleReq interface {
	HandleRequest()
}

// HandleRequest : starts server and handle html methods on routes
func HandleRequest() {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/breads", GetBreads).Methods("GET")
	Router.HandleFunc("/breads/{bread}", GetBread).Methods("GET")
	Router.HandleFunc("/breads", CreateBread).Methods("POST")
	Router.HandleFunc("/breads/{bread}", UpdateBread).Methods("PUT")
	Router.HandleFunc("/breads/{bread}", DeleteBread).Methods("DELETE")
	Router.HandleFunc("/breads", DeleteBreads).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
