package breadrecipe

// BreadRecipe : bread recipes model
type BreadRecipe struct {
	Name  string `json:"name" validate:"required"`
	Flour int    `json:"flour" validate:"required,numeric"`
	Water int    `json:"water" validate:"numeric"`
	Salt  int    `json:"salt" validate:"required,numeric"`
	Yeast int    `json:"yeast" validate:"required,numeric"`
	Sugar int    `json:"sugar" validate:"numeric"`
	Milk  int    `json:"milk" validate:"numeric"`
}
