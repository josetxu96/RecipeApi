package database

import (
	model "RecipeApi/internal/model/recipe"
	"database/sql"
)

// Store : database interface
type Store interface {
	CreateRecipe(recipe *model.Recipe) error
	GetRecipes() ([]*model.Recipe, error)
	GetRecipe(string) (model.Recipe, error)
	DeleteRecipe(name string) error
	DeleteRecipes() error
	UpdateRecipe(recipe *model.Recipe, name string) error
}

// DbStore : database struct
type DbStore struct {
	Db *sql.DB
}

// DB : database
var DB Store

// InitStore : starts the store
func InitStore(s Store) {
	DB = s
}
