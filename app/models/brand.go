package models

import "github.com/jinzhu/gorm"

type Brand struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(200)"`
}
