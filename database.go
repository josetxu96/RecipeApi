package mydatabase

import (
	"database/sql"
)

type Store interface {
	CreateBread(recipe *BreadRecipe) error
	GetBreads() ([]*BreadRecipe, error)
	GetBread(string) (BreadRecipe, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBread(recipe *BreadRecipe) error {

	_, err := store.db.Query("INSERT INTO breads(Name, Flour, Water, Salt, Yeast, Milk, Sugar) VALUES ($1,$2,$3,$4,$5,$6,$7)", recipe.Name, recipe.Flour, recipe.Water, recipe.Salt, recipe.Yeast, recipe.Milk, recipe.Sugar)
	return err
}

func (store *dbStore) GetBreads() ([]*BreadRecipe, error) {

	rows, err := store.db.Query("SELECT name, recipe FROM breads")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	breads := []*BreadRecipe{}
	for rows.Next() {

		bread := &BreadRecipe{}

		if err := rows.Scan(&bread.Name, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar); err != nil {
			return nil, err
		}

		breads = append(breads, bread)
	}
	return breads, nil
}

func (store *dbStore) GetBread(name string) (BreadRecipe, error) {

	row, err := store.db.Query("SELECT name, recipe FROM breads, WHERE name=?", name)

	if err != nil {
		return BreadRecipe{}, err
	}

	bread := BreadRecipe{}

	if err := row.Scan(&bread.Name, &bread.Flour, &bread.Water, &bread.Salt, &bread.Yeast, &bread.Milk, &bread.Sugar); err != nil {
		return BreadRecipe{}, err
	}

	return bread, nil
}

func InitStore(_connection *sql.DB) Store {
	return &dbStore{db: _connection}
}
