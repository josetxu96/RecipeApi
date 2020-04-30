package database

// DeleteRecipes : deletes all recipes from the database
func (store *DbStore) DeleteRecipes() error {
	_, err := store.Db.Exec("DELETE FROM recipes")
	return err
}
