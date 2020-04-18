package database

import (
	model "RecipeApi/internal/model/breadRecipe"
	"database/sql"
	"fmt"
)

type Store interface {
	CreateBread(recipe *model.BreadRecipe) error
	GetBreads() ([]*model.BreadRecipe, error)
	GetBread(string) (model.BreadRecipe, error)
	DeleteBread(name string) error
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBread(recipe *model.BreadRecipe) error {

	_, err := store.db.Query("INSERT INTO breads(name, flour, water, salt, yeast, milk, sugar) VALUES ($1,$2,$3,$4,$5,$6,$7)", recipe.Name, recipe.Flour, recipe.Water, recipe.Salt, recipe.Yeast, recipe.Milk, recipe.Sugar)
	return err
}

func (store *dbStore) GetBreads() ([]*model.BreadRecipe, error) {

	rows, err := store.db.Query("SELECT name, flour, water, salt, yeast, milk, sugar FROM breads")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	breads := []*model.BreadRecipe{}
	for rows.Next() {

		bread := &model.BreadRecipe{}

		if err := rows.Scan(&bread.Name, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar); err != nil {
			return nil, err
		}

		breads = append(breads, bread)
	}
	return breads, nil
}

func (store *dbStore) GetBread(name string) (model.BreadRecipe, error) {

	row := store.db.QueryRow("SELECT name, flour, water, salt, yeast, milk, sugar FROM breads where name = $1", name)

	bread := model.BreadRecipe{}
	err := row.Scan(&bread.Name, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar)
	if err != nil {
		return model.BreadRecipe{}, err
	}
	return bread, nil
}

func (store *dbStore) DeleteBread(name string) error {

	_, err := store.db.Exec("DELETE FROM breads where name = $1", name)
	fmt.Println("Deleting file")
	return err

}

func InitStore(db *sql.DB) Store {
	return &dbStore{db: db}
}
