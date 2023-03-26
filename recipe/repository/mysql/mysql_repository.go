package mysql

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/model"
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

func (mcp *mysqlRecipeRepository) GetById(id uint) (result *model.Recipe, err error) {
	cooker := &model.Cooker{
		ID: id,
	}

	fmt.Printf("%v", cooker)
	// var result model.Cooker{}

	// result := mcp.db.Model(cooker).First(cooker)
	// if result == mcp.db.Error {

	// }

	return result, nil
}

func (mcp *mysqlRecipeRepository) GetByCookerId(cookerId uint, pagination *model.Pagination) (*[]model.Recipe, error) {
	var results []model.Recipe

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := mcp.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Recipe{}).Where("cooker_id=?", cookerId).Find(&results)

	if result.Error != nil {
		message := result.Error
		return nil, message
	}

	return &results, nil
}

func (mcp *mysqlRecipeRepository) Store(*model.Recipe) (result *model.Recipe, err error) {
	return result, nil
}

func (mcp *mysqlRecipeRepository) Update(*model.Recipe) (result *model.Recipe, err error) {
	return result, nil
}
func (mcp *mysqlRecipeRepository) Delete(id uint) error {
	return nil
}
