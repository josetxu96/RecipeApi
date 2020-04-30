package database

import (
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
)

// CreateRecipe : adds a recipe to the database
func (store *DbStore) CreateRecipe(recipe *model.Recipe) error {
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Query("INSERT INTO recipes(name, description, ingredients, people, instructions) VALUES ($1,$2,$3,$4,$5)", recipe.Name, recipe.Description, s, recipe.People, recipe.Instructions)
	return err
}
