package Recipe

// Recipe : recipes model
type Recipe struct {
	Name         string         `json:"name" validate:"required"`
	Description  string         `json:"description" validate:"required"`
	People       int            `json:"people" validate:"required,numeric"`
	Ingredients  map[string]int `json:"ingredients" validate:"required"`
	Instructions string         `json:"instructions" validate:"required"`
}
