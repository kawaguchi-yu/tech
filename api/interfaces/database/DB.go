package database

import (
	"errors"
	"fmt"
	"hello/server/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) CreateUser(user domain.User) error { //渡された値をDBに入れる
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (db *UserRepository) CreateGood(good domain.Good) error { //渡された値をDBに入れる
	err := db.Create(&good).Error
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
func (db *UserRepository) GuestLogin() (u domain.User, err error) {
	var user domain.User
	if err := db.First(&user, "name = ?", "Guest User").Error; err != nil {
		fmt.Printf("ゲストログインに失敗しました\n")
		return user, err
	}
	return user, nil
}
func (db *UserRepository) ReturnUserBYEMail(email string) (u domain.User, err error) {
	var user domain.User
	if err := db.First(&user, "e_mail = ?", email).Error; err != nil {
		fmt.Printf("メールアドレスが存在しませんでした\n")
		return user, err
	}
	if err := db.First(&user.Profile, "user_id = ?", user.ID).Error; err != nil {
		fmt.Printf("profileがありませんでした\n")
		return user, nil
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
	if err := db.Where("user_id = ?", user.ID).Find(&posts).Error; err != nil {
		fmt.Printf("ポスト取得に失敗しました\n")
		return user, err
	}
	var goods []domain.Good
	if err := db.Find(&goods).Error; err != nil {
		fmt.Printf("いいねの取得に失敗しました\n")
		return user, err
	}
	if err := db.First(&user.Profile, "user_id", user.ID).Error; err != nil {
		fmt.Printf("profileの取得に失敗しました\n")
		return user, err
	}
	for _, post := range posts {
		for _, good := range goods {
			if post.ID == good.PostID {
				fmt.Printf("goodをpostにいれました%v\n", good.PostID)
				post.Goods = append(post.Goods, good)
			}
		}
		if post.UserID == user.ID {
			user.Posts = append(user.Posts, post)
		}
	}
	return user, nil
}
func (db *UserRepository) ReturnUserAndPostByPostID(id uint) (u domain.User, err error) {
	var post domain.Post
	var user domain.User
	var goods []domain.Good
	if err := db.First(&post, "id=?", id).Error; err != nil {
		fmt.Printf("ポスト取得に失敗しました%v\n", id)
		return user, err
	}
	if err := db.First(&user, "id=?", post.UserID).Error; err != nil {
		fmt.Printf("ユーザー取得に失敗しました%v\n", id)
		return user, err
	}
	if err := db.Find(&goods).Error; err != nil {
		fmt.Printf("いいね獲得に取得に失敗しました\n")
		return user, err
	}
	for _, good := range goods {
		if post.ID == good.PostID {
			fmt.Printf("goodをpostにいれました%v\n", good.PostID)
			post.Goods = append(post.Goods, good)
		}
	}
	user.Posts = append(user.Posts, post)
	return user, nil
}
func (db *UserRepository) DeleteGoodByPostID(postID uint, userID uint) (err error) {
	if err := db.Delete(domain.Good{}, "post_id=? AND user_id=?", postID, userID).Error; err != nil {
		fmt.Printf("deleteに失敗しましたpostid=%vuserid=%v\n", postID, userID)
		return err
	}
	fmt.Printf("正常に終了しましたpostid=%vuserid=%v\n", postID, userID)
	return nil
}
func (db *UserRepository) ReturnAllUserPost() (returnUsers []domain.User, returnPosts []domain.Post, returnGoods []domain.Good, err error) {
	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		return nil, nil, nil, err
	}
	var posts []domain.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, nil, nil, err
	}
	var goods []domain.Good
	if err := db.Find(&goods).Error; err != nil {
		return nil, nil, nil, err
	}
	return users, posts, goods, nil
}
func (db *UserRepository) ReturnGoodedPostByWord(word string) (returnUsers []domain.User, returnPosts []domain.Post, returnGoods []domain.Good, err error) {
	var users []domain.User
	var posts []domain.Post
	var goods []domain.Good
	var postId []uint
	if err := db.Find(&posts, "title LIKE ?", "%"+word+"%").Error; err != nil {
		return nil, nil, nil, err
	}
	for _, post := range posts {
		postId = append(postId, post.UserID)
	}
	if err := db.Find(&users, postId).Error; err != nil {
		return nil, nil, nil, err
	}
	if err := db.Find(&goods).Error; err != nil {
		return nil, nil, nil, err
	}
	return users, posts, goods, nil
}
func (db *UserRepository) ReturnGoodedPost(userID uint) (returnUsers []domain.User, returnPosts []domain.Post, returnGoods []domain.Good, err error) {
	var users []domain.User
	var posts []domain.Post
	var goods []domain.Good
	if err := db.Find(&goods, "user_id=?", userID).Error; err != nil {
		return nil, nil, nil, err
	}
	if len(goods) == 0 {
		return nil, nil, nil, err
	}
	var goodId []uint
	var postId []uint
	for _, good := range goods {
		goodId = append(goodId, good.PostID)
	}
	if err := db.Find(&posts, goodId).Error; err != nil {
		return nil, nil, nil, err
	}
	for _, post := range posts {
		postId = append(postId, post.UserID)
	}
	if err := db.Find(&users, postId).Error; err != nil {
		return nil, nil, nil, err
	}
	return users, posts, goods, nil
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
		return errors.New("編集権限がありません")
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
func (db *UserRepository) UpdateUser(email string, u domain.User) (err error) {
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return err
	}
	if u.Name != "" {
		user.Name = u.Name
		if err := db.Model(&user).Update("Name", u.Name).Error; err != nil {
			return err
		}
	}
	profile := new(domain.Profile)
	profile = &u.Profile
	fmt.Printf("%v\n", profile)
	if profile.ID == 0 {
		if err := db.Create(&profile).Error; err != nil {
			return err
		}
	} else {
		if err := db.Where("id = ?", profile.ID).Updates(&profile).Error; err != nil {
			return err
		}
	}
	return nil
}
