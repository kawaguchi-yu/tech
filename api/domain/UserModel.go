package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"Name" gorm:"unique"`
	EMail     string `json:"EMail" gorm:"unique"`
	Password  string `json:"Password"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	Profile   Profile
	ProfileID uint
	Goods     []Good
	Icon      string
}
