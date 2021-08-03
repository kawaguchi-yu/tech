package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID       uint
	Title        string `json:"Title"`
	Answer       string `json:"Answer"`
	WrongAnswer1 string `json:"WrongAnswer1"`
	WrongAnswer2 string `json:"WrongAnswer2"`
	WrongAnswer3 string `json:"WrongAnswer3"`
	Explanation  string `json:"Explanation"`
	Tags         []Tag
	Goods        []Good
}
