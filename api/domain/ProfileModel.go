package domain

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Essay  string
	Tags   []Tag
	URLs   []ExternaiURL
}
