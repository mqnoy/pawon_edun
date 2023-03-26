package domain

import (
	"example/web-service-gin/models"
)

type RecipeUsecase interface {
	CreateNewRecipe() error
	DetailRecipe(id int) (res models.Recipe, err error)
	UpdateRecipe() error
	DeleteRecipe() error
}

type RecipeRepository interface {
	TableName() string
	GetById(id uint) (*models.Recipe, error)
	GetByCookerId(cookerId uint, pagination *models.Pagination) (*[]models.Recipe, error)
	Store(*models.Recipe) (*models.Recipe, error)
	Update(*models.Recipe) (*models.Recipe, error)
	Delete(id uint) error
}
