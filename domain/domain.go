package domain

type Recipe struct {
	Id           int
	Name         string
	Ingredients  []string
	IsHalal      bool
	IsVegetarian bool
	Description  string
	Rating       float64
}

type RecipeRepository interface {
	FindById(id int) (*Recipe, error)
	Create(recipe Recipe) (*Recipe, error)
	Update(recipe Recipe) error
	Delete(id int) error
}
