package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(200)"`
	Cover       string `gorm:"column:cover;type:varchar(255)"`
	Description string `gorm:"column:description;type:text"`
	Slug        string `gorm:"column:slug;type:varchar(50)"`
	Status      int32    `gorm:"column:status;type:int"`
}
