package domain

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	PostID uint
	UserID uint
}
