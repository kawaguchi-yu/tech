package useCase

import "hello/server/domain"

type UserRepository interface {
	Store(domain.User) error
	SetIcon(string, string) error
	CreatePost(string, domain.Post) error
	DeletePost(string, domain.Post) error
	UpdatePost(string, *domain.Post) error
	DeleteAllByUserEMail(string) error
	ReturnUserBYEMail(string) (domain.User, error)
	ReturnUserPostByName(string) (domain.User, error)
	ReturnAllUserPost() ([]domain.User, []domain.Post, error)
}
