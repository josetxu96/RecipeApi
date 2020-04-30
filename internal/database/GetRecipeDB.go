package database

import (
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
)
// GetRecipe : gets a recipe from the database
func (store *DbStore) GetRecipe(name string) (model.Recipe, error) {

	row := store.Db.QueryRow("SELECT name, description, ingredients, people, instructions FROM recipes where name = $1", name)

	recipe := model.Recipe{}
	var s []byte
	err := row.Scan(&recipe.Name, &recipe.Description, &s, &recipe.People, &recipe.Instructions)
	if err != nil {
		return model.Recipe{}, err
	}
	json.Unmarshal(s, &recipe.Ingredients)
	return recipe, nil
}