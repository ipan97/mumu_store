package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"column:name;type:varchar(200)"`
	Price       float64 `gorm:"column:price;type:decimal(13,2)"`
	Quantity    int32   `gorm:"column:quantity;type:int"`
	Cover       string  `gorm:"column:cover;type:varchar(255)"`
	Description string  `gorm:"column:description;type:text"`
	Status      int32   `gorm:"column:status;type:int"`
}
