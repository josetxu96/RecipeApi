package database

import (
	model "RecipeApi/internal/model/breadrecipe"
	"database/sql"
	"encoding/json"
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
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Query("INSERT INTO breads(name, description, ingredients) VALUES ($1,$2,$3)", recipe.Name, recipe.Description, s)
	return err
}

// UpdateBread : updates a bread from the database
func (store *DbStore) UpdateBread(recipe *model.BreadRecipe, name string) error {
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Exec("UPDATE breads SET name=$1, description=$2, ingredients=$3 WHERE name=$4", recipe.Name, recipe.Description, s, name)
	return err
}

// GetBreads : gets all breads from the database
func (store *DbStore) GetBreads() ([]*model.BreadRecipe, error) {

	rows, err := store.Db.Query("SELECT name, description, ingredients FROM breads")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	breads := []*model.BreadRecipe{}
	for rows.Next() {

		bread := &model.BreadRecipe{}
		var s []byte
		if err := rows.Scan(&bread.Name, &bread.Description, &s); err != nil {
			return nil, err
		}
		json.Unmarshal(s, &bread.Ingredients)
		breads = append(breads, bread)
	}
	return breads, nil
}

// GetBread : gets a bread from the database
func (store *DbStore) GetBread(name string) (model.BreadRecipe, error) {

	row := store.Db.QueryRow("SELECT name, description, ingredients FROM breads where name = $1", name)

	bread := model.BreadRecipe{}
	var s []byte
	err := row.Scan(&bread.Name, &bread.Description, &s)
	if err != nil {
		return model.BreadRecipe{}, err
	}
	json.Unmarshal(s, &bread.Ingredients)
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
