package main

import (
	"RecipeApi/internal/database"
	"RecipeApi/internal/handler"
	"database/sql"

	_ "github.com/lib/pq"
)

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

	database.InitStore(&database.DbStore{Db: dbPostgres})
	handler.HandleRequest()
}