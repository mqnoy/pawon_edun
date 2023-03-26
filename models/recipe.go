package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID               uint            `json:"id" gorm:"primaryKey"`
	CookerID         uint            `json:"cooker_id" gorm:"column:cooker_id"`
	NumberOfServings int8            `json:"number_of_servings" gorm:"column:number_of_servings"`
	Title            string          `json:"title" gorm:"column:title"`
	Description      sql.NullString  `json:"description" gorm:"column:description"`
	Rating           sql.NullFloat64 `json:"rating" gorm:"column:rating"`
	Duration         int8            `json:"duration" gorm:"column:duration"`
}
