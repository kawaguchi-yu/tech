package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string
	ProfileID uint
	PostID    uint
}
