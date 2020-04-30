package database

import (
	model "RecipeApi/internal/model/recipe"
	"encoding/json"
)
// Getrecipes : gets all recipes from the database
func (store *DbStore) GetRecipes() ([]*model.Recipe, error) {

	rows, err := store.Db.Query("SELECT name, description, ingredients, people, instructions FROM recipes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []*model.Recipe{}
	for rows.Next() {

		recipe := &model.Recipe{}
		var s []byte
		if err := rows.Scan(&recipe.Name, &recipe.Description, &s, &recipe.People, &recipe.Instructions); err != nil {
			return nil, err
		}
		json.Unmarshal(s, &recipe.Ingredients)
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
