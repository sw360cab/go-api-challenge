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

	// proposal of extra fields to express and model a challenge
	/*
		TargetType string // Max,Min,Avg
		TargetUnit string // Meals, Steps,
		TargetAmount int // related to TargetUnit
		TargetPeriod customType // five nights in a row, a day for a month...
	*/
}
