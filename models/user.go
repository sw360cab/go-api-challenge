package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string
	Challenges []Challenge `gorm:"many2many:user_challenges;"`
}
