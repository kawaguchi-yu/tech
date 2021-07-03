package domain

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID   uint
	BirthDay time.Time
	Essay    string
	Tags     []Tag
	URLs     []ExternaiURL
}
