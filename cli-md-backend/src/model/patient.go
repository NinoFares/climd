package model

import "github.com/jinzhu/gorm"

type Patient struct {
	gorm.Model
	LastName  string
	FirstName string
	Age       int
	Number    string
}
