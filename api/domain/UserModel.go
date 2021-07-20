package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"Name"`
	EMail     string `json:"EMail" gorm:"unique"`
	Password  string `json:"Password"`
	Posts     []Post
	Profile   Profile
	ProfileID uint
	Goods     []Good
	icon      string
}
