package domain

import "gorm.io/gorm"

type ExternaiURL struct {
	gorm.Model
	Name      string
	URL       string
	ProfileID uint
}
