package domain

import (
	"example/web-service-gin/model"
)

type CookerUsecase interface {
	AddFavoriteRecipe() error
	RemoveFavoriteRecipe() error
	ListAllRecipes(cookerId uint, pagination *model.Pagination) (*[]model.Recipe, error)
}

type CookerRepository interface {
	TableName() string
	GetById(id uint) (*model.Cooker, error)
	GetByEmail(email string) (*model.Cooker, error)
	Store(*model.Cooker) (*model.Cooker, error)
	Update(*model.Cooker) (*model.Cooker, error)
	Delete(id uint) error
}
