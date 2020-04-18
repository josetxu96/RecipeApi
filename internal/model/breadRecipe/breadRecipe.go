package breadRecipe

type BreadRecipe struct {
	Name  string `json:"name" validate:"required"`
	Flour int    `json:"flour" validate:"required,numeric"`
	Water int    `json:"water" validate:"required,numeric"`
	Salt  int    `json:"salt" validate:"required,numeric"`
	Yeast int    `json:"yeast" validate:"required,numeric"`
	Sugar int    `json:"sugar" validate:"numeric"`
	Milk  int    `json:"milk" validate:"numeric"`
}
