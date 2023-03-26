package domain

import (
	"example/web-service-gin/model"
)

type RecipeUsecase interface {
	CreateNewRecipe() error
	DetailRecipe(id int) (res model.Recipe, err error)
	UpdateRecipe() error
	DeleteRecipe() error
}

type RecipeRepository interface {
	TableName() string
	GetById(id uint) (*model.Recipe, error)
	GetByCookerId(cookerId uint, pagination *model.Pagination) (*[]model.Recipe, error)
	Store(*model.Recipe) (*model.Recipe, error)
	Update(*model.Recipe) (*model.Recipe, error)
	Delete(id uint) error
}
