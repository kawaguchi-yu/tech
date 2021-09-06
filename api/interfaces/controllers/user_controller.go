package controllers

import (
	"fmt"
	"hello/server/domain"
	"hello/server/interfaces/database"
	"hello/server/useCase"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Interactor useCase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: useCase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}
func (controller *UserController) CreateUser(c Context) (err error) {
	user := domain.User{}
	c.Bind(&user)
	rawPassword := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, 4)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	fmt.Printf("%+v\n", user)
	user.Icon = ("dog_out.png")
	err = controller.Interactor.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ユーザー登録できませんでした")
	}
	claims := jwt.StandardClaims{
		Issuer:    user.EMail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encryptionKey, err := GetJWTToken()
	if err != nil {
		return err
	}
	tokenString, err := token.SignedString([]byte(encryptionKey)) //電子署名
	if err != nil {
		panic("電子署名できませんでした")
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(10 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	c.JSON(http.StatusOK, "Add成功！")
	return
}
func (controller *UserController) CreateGood(c Context) (err error) {
	good := domain.Good{}
	if err := c.Bind(&good); err != nil {
		fmt.Printf("Contextからuserを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "Contextからuserを読めませんでした")
	}
	err = controller.Interactor.CreateGood(good)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "goodできませんでした")
	}
	return c.JSON(http.StatusOK, "正常に終了しました")
}
func (controller *UserController) UpdateUser(c Context) (err error) {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		fmt.Printf("Contextからuserを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "Contextからuserを読めませんでした")
	}
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("Contextからemailを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "Contextからemailを読めませんでした")
	}
	err = controller.Interactor.UpdateUser(email, user)
	if err != nil {
		fmt.Printf("Contextからemailを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "updateできませんでした")
	}
	fmt.Printf("正常に終了しました\n")
	return c.JSON(http.StatusOK, "正常に終了しました")
}
func (controller *UserController) SetIcon(c Context) (err error) {
	icon, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("ファイルが読み込めません\n")
		return c.JSON(http.StatusBadRequest, icon)
	}
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, "クッキー読み取りに失敗しました")
	}
	src, err := icon.Open() //io.Readerに変換
	if err != nil {
		fmt.Printf("ファイルをioに変換できませんでした\n")
		return c.JSON(http.StatusBadRequest, "ファイルをioに変換できませんでした")
	}
	defer src.Close()

	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envファイルの読み込みが失敗しました\n")
	}
	awsAccesskey := os.Getenv("AWSACCESSKEY")
	awsSecretkey := os.Getenv("AWSSECRETKEY")
	fmt.Printf("key%v%v\n", awsAccesskey, awsSecretkey)
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsAccesskey, awsSecretkey, ""),
		Region:      aws.String("ap-northeast-1"),
	}))
	fmt.Printf("filename=%v\n", icon.Filename)
	uploader := s3manager.NewUploader(sess)
	// file, err := os.Open(icon.Filename)
	if err != nil {
		fmt.Printf("fileをopenできませんでした\n")
		return c.JSON(http.StatusBadRequest, "fileをopenできませんでした")
	}
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("techer-s3-001"),
		Key:    aws.String(icon.Filename),
		Body:   src,
	})
	if err != nil {
		fmt.Printf("uploadできませんでした errormessage=%v\n", res)
		return c.JSON(http.StatusBadRequest, "uploadできませんでした")
	}
	err = controller.Interactor.SetIcon(email, icon.Filename)
	if err != nil {
		fmt.Printf("iconをuserにセットできませんでした\n")
		return c.JSON(http.StatusBadRequest, "iconをuserにセットできませんでした")
	}
	fmt.Printf("seticonは正常に終了しました\n")
	return c.File("fallheal_out.png")
}
func (controller *UserController) CreatePost(c Context) (err error) {
	post := domain.Post{}
	c.Bind(&post)
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = controller.Interactor.CreatePost(email, post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, "正常に終了しました")
}
func (controller *UserController) DeletePost(c Context) (err error) {
	post := domain.Post{}
	c.Bind(&post)
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = controller.Interactor.DeletePost(email, post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, "正常に終了しました")
}
func (controller *UserController) UpdatePost(c Context) (err error) {
	post := new(domain.Post)
	c.Bind(post)
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = controller.Interactor.UpdatePost(email, post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, "正常に終了しました")
}

type Claims struct {
	jwt.StandardClaims
}

func (controller *UserController) DeleteUser(c Context) (err error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Printf("クッキーを読み取れませんでした%v\n", cookie)
		return c.JSON(500, "クッキーを読み取れませんでした")
	}
	encryptionKey, err := GetJWTToken()
	if err != nil {
		return err
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(encryptionKey), nil
	})
	if err != nil || !token.Valid {
		fmt.Printf("パルスに失敗しました\n")
		return c.JSON(http.StatusBadRequest, "パルスに失敗しました")
	}
	claims := token.Claims.(*Claims)
	email := claims.Issuer
	user := domain.User{}
	user, err = controller.Interactor.ReturnUserBYEMail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "emailからuserを取得できませんでした")
	}
	if err := GuestUserCheck(user.Name); err != nil {
		return c.JSON(http.StatusBadRequest, "GuestUserは削除権限がありません")
	}
	err = controller.Interactor.DeleteAllByUserEMail(email)
	if err != nil {
		c.JSON(500, "Delete失敗")
	}
	cookie.Expires = time.Now() //deleteしたuserのクッキーを消す
	c.SetCookie(cookie)
	fmt.Printf("ユーザーを削除しました。\n")
	c.JSON(http.StatusOK, "Delete成功！")
	return
}
func (controller *UserController) DeleteGoodByPostID(c Context) (err error) {
	good := new(domain.Good)
	c.Bind(good)
	if err := controller.Interactor.DeleteGoodByPostID(good.PostID, good.UserID); err != nil {
		return c.JSON(http.StatusBadRequest, "Goodを削除できませんでした")
	}
	return c.JSON(http.StatusOK, "正常に終了しました")
}
func (controller *UserController) GuestLogin(c Context) (err error) {
	user, err := controller.Interactor.GuestLogin()
	if err != nil {
		fmt.Printf("メールアドレスが存在しませんでした\n")
		return c.JSON(http.StatusBadRequest, "メールアドレスが存在しませんでした")
	}
	claims := jwt.StandardClaims{
		Issuer:    user.EMail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encryptionKey, err := GetJWTToken()
	if err != nil {
		return err
	}
	tokenString, err := token.SignedString([]byte(encryptionKey)) //電子署名
	if err != nil {
		panic("電子署名できませんでした")
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(10 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, cookie)
}
func (controller *UserController) Login(c Context) (err error) {
	u := new(domain.User)
	c.Bind(u)
	user, err := controller.Interactor.ReturnUserBYEMail(u.EMail)
	if err != nil {
		fmt.Printf("メールアドレスが存在しませんでした\n")
		return c.JSON(http.StatusBadRequest, "メールアドレスが存在しませんでした")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, "パスワードが違います")
	}
	claims := jwt.StandardClaims{
		Issuer:    u.EMail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encryptionKey, err := GetJWTToken()
	if err != nil {
		return err
	}
	tokenString, err := token.SignedString([]byte(encryptionKey)) //電子署名
	if err != nil {
		panic("電子署名できませんでした")
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(10 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, cookie)
}
func (controller *UserController) Logout(c Context) (err error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, cookie)
	}
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	fmt.Printf("Logoutに成功しました\n")
	return c.JSON(http.StatusOK, cookie)
}
func (controller *UserController) ReadCookieReturnUser(c Context) (err error) {
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		return err
	}
	user, err := controller.Interactor.ReturnUserBYEMail(email)
	if err != nil {
		return err
	}
	user.Icon = getIcon(user.Icon)
	fmt.Printf("user.Name=%v\n", user.Name)
	return c.JSON(http.StatusOK, user)
}
func (controller *UserController) ReturnAllUserPost(c Context) error {
	users, posts, goods, err := controller.Interactor.ReturnAllUserPost()
	if err != nil {
		fmt.Printf("ユーザーを取得できませんでした\n")
		return c.JSON(http.StatusBadRequest, "ユーザーを取得できませんでした")
	}
	var returnUsers []domain.User
	for _, user := range users {
		for _, post := range posts {
			for _, good := range goods {
				if post.ID == good.PostID {
					post.Goods = append(post.Goods, good)
				}
			}
			if post.UserID == user.ID {
				user.Posts = append(user.Posts, post)
			}
		}
		if user.Posts != nil {
			user.Icon = getIcon(user.Icon)
			returnUsers = append(returnUsers, user)
		}
	}
	fmt.Printf("ReturnAllUserは正常に終了しました\n")
	return c.JSON(http.StatusOK, returnUsers)
}
func (controller *UserController) ReturnUserPostByName(c Context) error {
	u := new(domain.User)
	c.Bind(u)
	user, err := controller.Interactor.ReturnUserPostByName(u.Name)
	if err != nil {
		c.JSON(http.StatusOK, "名前から記事を取得できませんでした")
	}
	user.Icon = getIcon(user.Icon)
	if user.Icon == "" {
		return c.JSON(http.StatusOK, "base64エンコードに失敗しました")
	}
	fmt.Printf("ReturnUserPostByNameは正常に終了しました\n")
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) ReturnUserAndPostByPostID(c Context) error {
	post := new(domain.Post)
	c.Bind(post)
	user, err := controller.Interactor.ReturnUserAndPostByPostID(post.UserID)
	if err != nil {
		c.JSON(http.StatusOK, "idからユーザーを取得できませんでした")
	}
	return c.JSON(http.StatusOK, user)
}

type search struct {
	Word string
}

func (controller *UserController) ReturnGoodedPostByWord(c Context) error {
	data := new(search)
	if err := c.Bind(data); err != nil {
		return err
	}
	fmt.Printf("%v\n", data)
	if data.Word == "" {
		return c.JSON(http.StatusBadRequest, "wordがnullです")
	}
	users, posts, goods, err := controller.Interactor.ReturnGoodedPostByWord(data.Word)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ユーザー情報が取得できませんでした")
	}
	var returnUsers []domain.User
	for _, user := range users {
		for _, post := range posts {
			for _, good := range goods {
				if post.ID == good.PostID {
					post.Goods = append(post.Goods, good)
				}
			}
			if user.ID == post.UserID {
				user.Posts = append(user.Posts, post)
			}
		}
		if user.Posts != nil {
			user.Icon = getIcon(user.Icon)
			returnUsers = append(returnUsers, user)
		}
	}
	fmt.Printf("正常に終了しました\n")
	return c.JSON(http.StatusOK, returnUsers)
}
func (controller *UserController) ReturnGoodedPost(c Context) error {
	good := new(domain.Good)
	c.Bind(good)
	fmt.Printf("%v\n", good)
	users, posts, goods, err := controller.Interactor.ReturnGoodedPost(good.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "記事を取得できませんでした")
	}
	var returnUsers []domain.User
	for _, user := range users {
		for _, post := range posts {
			for _, good := range goods {
				if post.ID == good.PostID {
					post.Goods = append(post.Goods, good)
				}
			}
			if user.ID == post.UserID {
				user.Posts = append(user.Posts, post)
			}
		}
		if user.Posts != nil {
			user.Icon = getIcon(user.Icon)
			returnUsers = append(returnUsers, user)
		}
	}
	fmt.Printf("正常に終了しました\n")
	return c.JSON(http.StatusOK, returnUsers)
}
