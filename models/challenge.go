package models

import (
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model  `json:"-"`
	Name        string
	Description string
	Available   bool   `json:"-" gorm:"default:true"`
	Users       []User `json:"-" gorm:"many2many:user_challenges;"`
	// cmp. https://gorm.cn/docs/create.html#Default-Values
}
