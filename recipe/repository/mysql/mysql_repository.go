package mysql

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/models"
	"fmt"

	"gorm.io/gorm"
)

type mysqlRecipeRepository struct {
	DB *gorm.DB
}

func NewMysqlRecipeRepository(db *gorm.DB) domain.RecipeRepository {
	return &mysqlRecipeRepository{
		DB: db,
	}
}

func (mcp *mysqlRecipeRepository) TableName() string {
	return "recipes"
}

func (mcp *mysqlRecipeRepository) GetById(id uint) (result *models.Recipe, err error) {
	cooker := &models.Cooker{
		ID: id,
	}

	fmt.Printf("%v", cooker)
	// var result models.Cooker{}

	// result := mcp.db.Model(cooker).First(cooker)
	// if result == mcp.db.Error {

	// }

	return result, nil
}

func (mcp *mysqlRecipeRepository) GetByCookerId(cookerId uint, pagination *models.Pagination) (*[]models.Recipe, error) {
	var results []models.Recipe

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := mcp.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Recipe{}).Where("cooker_id=?", cookerId).Find(&results)

	if result.Error != nil {
		message := result.Error
		return nil, message
	}

	return &results, nil
}

func (mcp *mysqlRecipeRepository) Store(*models.Recipe) (result *models.Recipe, err error) {
	return result, nil
}

func (mcp *mysqlRecipeRepository) Update(*models.Recipe) (result *models.Recipe, err error) {
	return result, nil
}
func (mcp *mysqlRecipeRepository) Delete(id uint) error {
	return nil
}
