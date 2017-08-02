package model

import "github.com/jinzhu/gorm"

type User struct {
}

type TestAutoMigration struct {
	gorm.Model
	Name     string
	Interest string `gorm:"size:250"`
	Age      int
}
