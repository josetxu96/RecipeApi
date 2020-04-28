package breadrecipe

// BreadRecipe : bread recipes model
type BreadRecipe struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	People      int            `json:"people" validate:"required"`
	Ingredients map[string]int `json:"ingredients" validate:"required"`
}
