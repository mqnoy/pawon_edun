package mysql

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/models"
	"fmt"

	"gorm.io/gorm"
)

type mysqlCookerRepository struct {
	DB *gorm.DB
}

func NewMysqlCookerRepository(db *gorm.DB) domain.CookerRepository {
	return &mysqlCookerRepository{
		DB: db,
	}
}

func (mcp *mysqlCookerRepository) TableName() string {
	return "cookers"
}

func (mcp *mysqlCookerRepository) GetById(id uint) (result *models.Cooker, err error) {
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

func (mcp *mysqlCookerRepository) GetByEmail(email string) (result *models.Cooker, err error) {
	return result, nil
}

func (mcp *mysqlCookerRepository) Store(*models.Cooker) (result *models.Cooker, err error) {
	return result, nil
}

func (mcp *mysqlCookerRepository) Update(*models.Cooker) (result *models.Cooker, error error) {
	return result, nil
}
func (mcp *mysqlCookerRepository) Delete(id uint) error {
	return nil
}
