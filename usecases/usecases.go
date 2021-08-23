package usecases

import (
	"mse-publisher/domain"
)

type Recipe struct {
	Id           int
	Name         string
	Ingredients  []string
	IsHalal      bool
	IsVegetarian bool
	Description  string
	Rating       float64
}

type RecipeInteractor struct {
	RecipeRepository domain.RecipeRepository
}

func (interactor *RecipeInteractor) GetRecipeById(recipeId int) (res *Recipe, err error) {
	recipe, err := interactor.RecipeRepository.FindById(recipeId)
	if err != nil {
		return nil, err
	}
	res = &Recipe{
		Id:           recipe.Id,
		Name:         recipe.Name,
		Ingredients:  recipe.Ingredients,
		IsHalal:      recipe.IsHalal,
		IsVegetarian: recipe.IsVegetarian,
		Description:  recipe.Description,
		Rating:       recipe.Rating,
	}
	return res, nil
}

func (interactor *RecipeInteractor) InsertRecipe(recipe Recipe) (*Recipe, error) {
	rec, err := interactor.RecipeRepository.Create(domain.Recipe{
		Name:         recipe.Name,
		Ingredients:  recipe.Ingredients,
		IsHalal:      recipe.IsHalal,
		IsVegetarian: recipe.IsVegetarian,
		Description:  recipe.Description,
		Rating:       recipe.Rating,
	})
	if err != nil {
		return nil, err
	}
	return &Recipe{
		Id:           rec.Id,
		Name:         rec.Name,
		Ingredients:  rec.Ingredients,
		IsHalal:      rec.IsHalal,
		IsVegetarian: rec.IsVegetarian,
		Description:  rec.Description,
		Rating:       rec.Rating,
	}, nil
}

func (interactor *RecipeInteractor) UpdateRecipe(recipe Recipe) error {
	err := interactor.RecipeRepository.Update(domain.Recipe{
		Id:           recipe.Id,
		Name:         recipe.Name,
		Ingredients:  recipe.Ingredients,
		IsHalal:      recipe.IsHalal,
		IsVegetarian: recipe.IsVegetarian,
		Description:  recipe.Description,
		Rating:       recipe.Rating,
	})
	if err != nil {
		return err
	}
	return nil
}

func (interactor *RecipeInteractor) DeleteRecipe(id int) error {
	err := interactor.RecipeRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
