package usecase

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/model"

	"github.com/sirupsen/logrus"
)

type cookerUsecase struct {
	cookerRepo domain.CookerRepository
	recipeRepo domain.RecipeRepository
}

func NewCookerUsecase(cr domain.CookerRepository, rr domain.RecipeRepository) domain.CookerUsecase {
	return &cookerUsecase{
		cookerRepo: cr,
		recipeRepo: rr,
	}
}

func (cu *cookerUsecase) AddFavoriteRecipe() error {
	return nil
}

func (cu *cookerUsecase) RemoveFavoriteRecipe() error {
	return nil
}

func (cu *cookerUsecase) ListAllRecipes(cookerId uint, pagination *model.Pagination) (*[]model.Recipe, error) {
	recipes, err := cu.recipeRepo.GetByCookerId(cookerId, pagination)

	if err != nil {
		logrus.Error(err)
	}

	return recipes, nil
}
