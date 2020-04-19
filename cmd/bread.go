package main

import (
	"RecipeApi/internal/database"
	"RecipeApi/internal/handler"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func handleRequest(h handler.Handler) {

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/breads", h.GetBreads).Methods("GET")
	Router.HandleFunc("/breads/{bread}", h.GetBread).Methods("GET")
	Router.HandleFunc("/breads", h.CreateBread).Methods("POST")
	Router.HandleFunc("/breads/{bread}", h.UpdateBread).Methods("PUT")
	Router.HandleFunc("/breads/{bread}", h.DeleteBread).Methods("DELETE")
	Router.HandleFunc("/breads", h.DeleteBreads).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func main() {
	connString := "user=jose password=42771618210 dbname=breads sslmode=disable"
	dbPostgres, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = dbPostgres.Ping()

	if err != nil {
		panic(err)
	}

	db := database.InitStore(dbPostgres)
	h := handler.InitHandler(db)
	handleRequest(h)

}

//docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=42771618210 -e POSTGRES_USER=jose -d postgres
