package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Tags    string `gorm:"type:text" json:"tags"`
}
