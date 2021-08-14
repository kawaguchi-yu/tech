package useCase

import (
	"hello/server/domain"
)

type UserInteractor struct {
	UserRepository UserRepository //user_repositoryのinterfaceで中身を定義
}

func (interactor *UserInteractor) CreateUser(user domain.User) (err error) {
	err = interactor.UserRepository.CreateUser(user)
	return
}
func (interactor *UserInteractor) CreateGood(good domain.Good) (err error) {
	err = interactor.UserRepository.CreateGood(good)
	return
}
func (interactor *UserInteractor) SetIcon(email string, IconPath string) (err error) {
	err = interactor.UserRepository.SetIcon(email, IconPath)
	return
}
func (interactor *UserInteractor) DeleteAllByUserEMail(email string) (err error) {
	err = interactor.UserRepository.DeleteAllByUserEMail(email)
	return
}
func (interactor *UserInteractor) DeleteGoodByPostID(postID uint, userID uint) (err error) {
	err = interactor.UserRepository.DeleteGoodByPostID(postID, userID)
	return
}
func (interactor *UserInteractor) ReturnUserBYEMail(email string) (user domain.User, err error) {
	user, err = interactor.UserRepository.ReturnUserBYEMail(email)
	return
}
func (interactor *UserInteractor) ReturnAllUserPost() (users []domain.User, posts []domain.Post, err error) {
	users, posts, err = interactor.UserRepository.ReturnAllUserPost()
	return
}
func (interactor *UserInteractor) ReturnUserPostByName(name string) (user domain.User, err error) {
	user, err = interactor.UserRepository.ReturnUserPostByName(name)
	return
}
func (interactor *UserInteractor) ReturnGoodedPost(userID uint) (users []domain.User, posts []domain.Post, goods []domain.Good, err error) {
	users, posts, goods, err = interactor.UserRepository.ReturnGoodedPost(userID)
	return
}
func (interactor *UserInteractor) ReturnUserAndPostByPostID(id uint) (user domain.User, err error) {
	user, err = interactor.UserRepository.ReturnUserAndPostByPostID(id)
	return
}
func (interactor *UserInteractor) CreatePost(email string, post domain.Post) (err error) {
	err = interactor.UserRepository.CreatePost(email, post)
	return
}
func (interactor *UserInteractor) GuestLogin() (user domain.User, err error) {
	user, err = interactor.UserRepository.GuestLogin()
	return
}
func (interactor *UserInteractor) DeletePost(email string, post domain.Post) (err error) {
	err = interactor.UserRepository.DeletePost(email, post)
	return
}
func (interactor *UserInteractor) UpdatePost(email string, post *domain.Post) (err error) {
	err = interactor.UserRepository.UpdatePost(email, post)
	return
}
func (interactor *UserInteractor) UpdateUser(email string, u domain.User) (err error) {
	err = interactor.UserRepository.UpdateUser(email, u)
	return
}
