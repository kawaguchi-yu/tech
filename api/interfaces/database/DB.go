package database

import (
	"fmt"
	"hello/server/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(user domain.User) error { //渡された値をDBに入れる
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (db *UserRepository) DeleteAllByUserEMail(email string) error {
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	if err := db.Delete(&domain.User{}, user.ID).Error; err != nil {
		return err
	}
	if err := db.Delete(&domain.Post{}, "user_id = ?", user.ID).Error; err != nil {
		return err
	}
	return nil
}
func (db *UserRepository) DeletePost(email string, post domain.Post) error {
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	if user.ID == post.UserID {
		db.Delete(&domain.Post{}, post.ID)
	} else {
		fmt.Printf("削除権限がありません\n")
		return nil
	}
	return nil
}
func (db *UserRepository) ReturnUserBYEMail(email string) (u domain.User, err error) {
	var user domain.User
	if err := db.First(&user, "e_mail = ?", email).Error; err != nil {
		fmt.Printf("メールアドレスが存在しませんでした\n")
		return user, err
	}
	return user, nil
}
func (db *UserRepository) ReturnUserPostByName(name string) (u domain.User, err error) {
	var user domain.User
	if err := db.First(&user, "name=?", name).Error; err != nil {
		fmt.Printf("ユーザー取得に失敗しました%v\n", name)
		return user, err
	}
	var posts []domain.Post
	db.Where("user_id = ?", user.ID).Find(&posts)
	for _, post := range posts {
		if post.UserID == user.ID {
			user.Posts = append(user.Posts, post)
		}
	}
	return user, nil
}
func (db *UserRepository) ReturnAllUserPost() (users []domain.User, posts []domain.Post, err error) {
	var user []domain.User
	if err := db.Find(&user).Error; err != nil {
		return nil, nil, err
	}
	var post []domain.Post
	if err := db.Find(&post).Error; err != nil {
		return nil, nil, err
	}
	return user, post, nil
}
func (db *UserRepository) CreatePost(email string, post domain.Post) (err error) {
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	post.UserID = user.ID
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}
func (db *UserRepository) UpdatePost(email string, post *domain.Post) (err error) {
	fmt.Printf("postID=%v\n", post.ID)
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	if user.ID == post.UserID {
		if err := db.Model(&domain.Post{}).Where("Id =?", post.ID).Updates(post).Error; err != nil {
			fmt.Printf("アップデートに失敗しました\n")
			return err
		}
		fmt.Printf("postの%vをアップデートしました\n", post.ID)
	} else {
		fmt.Printf("編集権限がありませんuser.ID=%vpost.UserID=%v\n", user.ID, post.UserID)
		return nil
	}
	return nil
}
func (db *UserRepository) SetIcon(email string, IconPath string) (err error) {
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	if err := db.Model(&user).Update("icon", IconPath).Error; err != nil {
		return err
	}
	return nil
}
