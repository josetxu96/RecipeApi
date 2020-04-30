package database

// DeleteRecipe : deletes a recipe from the database
func (store *DbStore) DeleteRecipe(name string) error {

	_, err := store.Db.Exec("DELETE FROM recipes where name = $1", name)
	return err

}
