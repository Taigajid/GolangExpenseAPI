package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	UserID  int
	Expense string
	Price   float64
}
