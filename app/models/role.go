package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string
	DisplayName string
	Description string
}
