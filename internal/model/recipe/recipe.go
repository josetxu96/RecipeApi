package Recipe

// Recipe : recipes model
type Recipe struct {
	Name         string               `json:"name" validate:"required"`
	Description  string               `json:"description" validate:"required"`
	People       int                  `json:"people" validate:"required,numeric"`
	Ingredients  map[string]*Ingredient `json:"ingredients" validate:"required"`
	Instructions string               `json:"instructions" validate:"required"`
}

type Ingredient struct {
	Quantity int    `json:"quantity" validate:"required,numeric"`
	Measure  string `json:"measure" validate:"required"`
}
