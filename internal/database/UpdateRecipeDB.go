package database

import (
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
)

// UpdateRecipe : updates a recipe from the database
func (store *DbStore) UpdateRecipe(recipe *model.Recipe, name string) error {
	s, _ := json.Marshal(recipe.Ingredients)
	_, err := store.Db.Exec("UPDATE recipes SET name=$1, description=$2, ingredients=$3, people=$4, instructions=$5 WHERE name=$6", recipe.Name, recipe.Description, s, recipe.People, recipe.Instructions, name)
	return err
}