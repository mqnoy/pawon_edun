package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Cooker struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"column:email"`
	Name      sql.NullString `json:"name" gorm:"column:name"`
	Password  string         `json:"password" gorm:"column:password"`
	LastLogin sql.NullTime   `json:"last_login" gorm:"column:last_login"`
	IsActive  int8           `json:"is_active" gorm:"column:is_active"`
	Recipes   []Recipe       `json:"recipes"`
}
