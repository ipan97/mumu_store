package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username    string  `gorm:"column:username"`
	Password    string  `gorm:"column:password"`
	Email       *string `gorm:"column:email"`
	PhoneNumber *string `gorm:"column:phone_number"`
	Address     string  `gorm:"column:address"`
}
