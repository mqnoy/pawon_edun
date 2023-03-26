package usecase

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/models"
)

type recipeUsecase struct {
	recipeRepo domain.RecipeRepository
	cookerRepo domain.CookerRepository
}

func NewRecipeUsecase(rr domain.RecipeRepository, cr domain.CookerRepository) domain.RecipeUsecase {
	return &recipeUsecase{
		recipeRepo: rr,
		cookerRepo: cr,
	}
}

func (ru *recipeUsecase) CreateNewRecipe() error {
	return nil
}

func (ru *recipeUsecase) DetailRecipe(id int) (res models.Recipe, err error) {
	// find recipe
	// resRecipe, err   =ru.recipeRepo.GetById(id)
	// if err != nill{
	// 	return
	// }

	// find cooker
	// resCooker, err = ru.coo

	return res, nil
}

func (ru *recipeUsecase) UpdateRecipe() error {
	return nil
}

func (ru *recipeUsecase) DeleteRecipe() error {
	return nil
}
