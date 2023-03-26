package mysql

import (
	"example/web-service-gin/domain"
	"example/web-service-gin/model"
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

func (mcp *mysqlCookerRepository) GetById(id uint) (result *model.Cooker, err error) {
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

func (mcp *mysqlCookerRepository) GetByEmail(email string) (result *model.Cooker, err error) {
	return result, nil
}

func (mcp *mysqlCookerRepository) Store(*model.Cooker) (result *model.Cooker, err error) {
	return result, nil
}

func (mcp *mysqlCookerRepository) Update(*model.Cooker) (result *model.Cooker, error error) {
	return result, nil
}
func (mcp *mysqlCookerRepository) Delete(id uint) error {
	return nil
}
