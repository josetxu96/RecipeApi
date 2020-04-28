package database

import (
	model "RecipeApi/internal/model/recipe"
	"database/sql"
	"encoding/json"
)

// Store : database interface
type Store interface {
	CreateRecipe(recipe *model.Recipe) error
	Getrecipes() ([]*model.Recipe, error)
	Getrecipe(string) (model.Recipe, error)
	Deleterecipe(name string) error
	Deleterecipes() error
	Updaterecipe(recipe *model.Recipe, name string) error
}

// DbStore : database struct
type DbStore struct {
	Db *sql.DB
}

// DB : database
var DB Store

// Createrecipe : adds a recipe to the database
func (store *DbStore) CreateRecipe(recipe *model.Recipe) error {
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Query("INSERT INTO recipes(name, description, ingredients, people) VALUES ($1,$2,$3,$4)", recipe.Name, recipe.Description, s, recipe.People)
	return err
}

// Updaterecipe : updates a recipe from the database
func (store *DbStore) UpdateRecipe(recipe *model.Recipe, name string) error {
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Exec("UPDATE recipes SET name=$1, description=$2, ingredients=$3, people=$4 WHERE name=$5", recipe.Name, recipe.Description, s, recipe.People, name)
	return err
}

// Getrecipes : gets all recipes from the database
func (store *DbStore) GetRecipes() ([]*model.Recipe, error) {

	rows, err := store.Db.Query("SELECT name, description, ingredients, people FROM recipes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []*model.Recipe{}
	for rows.Next() {

		recipe := &model.Recipe{}
		var s []byte
		if err := rows.Scan(&recipe.Name, &recipe.Description, &s, &recipe.People); err != nil {
			return nil, err
		}
		json.Unmarshal(s, &recipe.Ingredients)
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

// Getrecipe : gets a recipe from the database
func (store *DbStore) GetRecipe(name string) (model.Recipe, error) {

	row := store.Db.QueryRow("SELECT name, description, ingredients, people FROM recipes where name = $1", name)

	recipe := model.Recipe{}
	var s []byte
	err := row.Scan(&recipe.Name, &recipe.Description, &s, &recipe.People)
	if err != nil {
		return model.Recipe{}, err
	}
	json.Unmarshal(s, &recipe.Ingredients)
	return recipe, nil
}

// Deleterecipe : deletes a recipe from the database
func (store *DbStore) DeleteRecipe(name string) error {

	_, err := store.Db.Exec("DELETE FROM recipes where name = $1", name)
	return err

}

// Deleterecipes : deletes all recipes from the database
func (store *DbStore) DeleteRecipes() error {
	_, err := store.Db.Exec("DELETE FROM recipes")
	return err
}

// InitStore : starts the store
func InitStore(s Store) {
	DB = s
}
