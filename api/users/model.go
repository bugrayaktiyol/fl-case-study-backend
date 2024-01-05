package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Age     int
	Address string
	Tags    []string `gorm:"type:text"`
}
