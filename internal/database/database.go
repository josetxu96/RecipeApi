package database

import (
	model "RecipeApi/internal/model/breadrecipe"
	"database/sql"
)

// Store : database interface
type Store interface {
	CreateBread(recipe *model.BreadRecipe) error
	GetBreads() ([]*model.BreadRecipe, error)
	GetBread(string) (model.BreadRecipe, error)
	DeleteBread(name string) error
	DeleteBreads() error
	UpdateBread(recipe *model.BreadRecipe, name string) error
}

// DbStore : database struct
type DbStore struct {
	Db *sql.DB
}

// DB : database
var DB Store

// CreateBread : adds a bread to the database
func (store *DbStore) CreateBread(recipe *model.BreadRecipe) error {

	_, err := store.Db.Query("INSERT INTO breads(name, description, flour, water, salt, yeast, milk, sugar) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", recipe.Name, recipe.Description, recipe.Flour, recipe.Water, recipe.Salt, recipe.Yeast, recipe.Milk, recipe.Sugar)
	return err
}

// UpdateBread : updates a bread from the database
func (store *DbStore) UpdateBread(recipe *model.BreadRecipe, name string) error {
	_, err := store.Db.Exec("UPDATE breads SET name=$1, description=$2, flour=$3, water=$4, salt=$5, yeast=$6, milk=$7, sugar=$8 WHERE name=$9", recipe.Name, recipe.Description, recipe.Flour, recipe.Water, recipe.Salt, recipe.Yeast, recipe.Milk, recipe.Sugar, name)
	return err
}

// GetBreads : gets all breads from the database
func (store *DbStore) GetBreads() ([]*model.BreadRecipe, error) {

	rows, err := store.Db.Query("SELECT name, description, flour, water, salt, yeast, milk, sugar FROM breads")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	breads := []*model.BreadRecipe{}
	for rows.Next() {

		bread := &model.BreadRecipe{}

		if err := rows.Scan(&bread.Name, &bread.Description, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar); err != nil {
			return nil, err
		}

		breads = append(breads, bread)
	}
	return breads, nil
}

// GetBread : gets a bread from the database
func (store *DbStore) GetBread(name string) (model.BreadRecipe, error) {

	row := store.Db.QueryRow("SELECT name, description, flour, water, salt, yeast, milk, sugar FROM breads where name = $1", name)

	bread := model.BreadRecipe{}
	err := row.Scan(&bread.Name, &bread.Description, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar)
	if err != nil {
		return model.BreadRecipe{}, err
	}
	return bread, nil
}

// DeleteBread : deletes a bread from the database
func (store *DbStore) DeleteBread(name string) error {

	_, err := store.Db.Exec("DELETE FROM breads where name = $1", name)
	return err

}

// DeleteBreads : deletes all breads from the database
func (store *DbStore) DeleteBreads() error {
	_, err := store.Db.Exec("DELETE FROM breads")
	return err
}

// InitStore : starts the store
func InitStore(s Store) {
	DB = s
}
