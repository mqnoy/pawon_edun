package domain

import (
	"example/web-service-gin/models"
)

type CookerUsecase interface {
	AddFavoriteRecipe() error
	RemoveFavoriteRecipe() error
	ListAllRecipes(cookerId uint, pagination *models.Pagination) (*[]models.Recipe, error)
}

type CookerRepository interface {
	TableName() string
	GetById(id uint) (*models.Cooker, error)
	GetByEmail(email string) (*models.Cooker, error)
	Store(*models.Cooker) (*models.Cooker, error)
	Update(*models.Cooker) (*models.Cooker, error)
	Delete(id uint) error
}
