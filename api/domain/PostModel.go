package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID uint
	Name   string
	Text   string
	Tags   []Tag
	Goods  []Good
}
