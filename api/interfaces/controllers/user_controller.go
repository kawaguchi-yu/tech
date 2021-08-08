package controllers

import (
	"fmt"
	"hello/server/domain"
	"hello/server/infra/useCase"
	"hello/server/interfaces/database"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	u := domain.User{}
	c.Bind(&u)
	rawPassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, 4)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	fmt.Printf("%+v\n", u)
	u.Icon = ("dog_out.png")
	err = controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, "Add失敗")
	}
	claims := jwt.StandardClaims{
		Issuer:    u.EMail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("kirikiri")) //電子署名
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
	c.JSON(201, "Add成功！")
	return
}
func (controller *UserController) UpdateUser(c Context) (err error) {
	u := domain.User{}
	if err := c.Bind(&u); err != nil {
		fmt.Printf("Contextからuserを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "Contextからuserを読めませんでした")
	}
	fmt.Printf("user=%v\n", u)
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("Contextからemailを読めませんでした\n")
		c.JSON(http.StatusBadRequest, "Contextからemailを読めませんでした")
	}
	err = controller.Interactor.UpdateUser(email, u)
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
	src, err := icon.Open() //io.Readerに変換
	if err != nil {
		fmt.Printf("ファイルをioに変換できませんでした\n")
		return c.JSON(http.StatusBadRequest, "ファイルをioに変換できませんでした")
	}
	defer src.Close()
	os.Chdir("img")
	iconModel := strings.Split(icon.Filename, ".")
	iconName := iconModel[0]
	extension := iconModel[1]
	dst, err := os.Create(fmt.Sprintf("%s_out.%s", iconName, extension))
	if err != nil { //"%s_out.%s"ここに\nを付け足すな！！！！！
		fmt.Printf("ファイルが作れませんでした\n")
		return c.JSON(http.StatusBadRequest, "ファイルが作れませんでした")
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil { //ファイルの内容をコピー
		fmt.Printf("コピーできませんでした\n")
		return c.JSON(http.StatusBadRequest, "コピーできませんでした")
	}
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, "クッキー読み取りに失敗しました")
	}
	err = controller.Interactor.SetIcon(email, dst.Name())
	if err != nil {
		fmt.Printf("iconをuserにセットできませんでした\n")
		return c.JSON(http.StatusBadRequest, "iconをuserにセットできませんでした")
	}
	fmt.Printf("seticonは正常に終了しました\n")
	return c.File(dst.Name())
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
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("kirikiri"), nil
	})
	if err != nil || !token.Valid {
		fmt.Printf("パルスに失敗しました\n")
		return c.JSON(500, "パルスに失敗しました")
	}
	claims := token.Claims.(*Claims)
	email := claims.Issuer
	err = controller.Interactor.DeleteAllByUserEMail(email)
	if err != nil {
		c.JSON(500, "Delete失敗")
	}
	cookie.Expires = time.Now() //deleteしたuserのクッキーを消す
	c.SetCookie(cookie)
	fmt.Printf("ユーザーを削除しました。\n")
	c.JSON(201, "Delete成功！")
	return
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
	tokenString, err := token.SignedString([]byte("kirikiri")) //電子署名
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
	tokenString, err := token.SignedString([]byte("kirikiri")) //電子署名
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
	users, posts, err := controller.Interactor.ReturnAllUserPost()
	if err != nil {
		fmt.Printf("ユーザーを取得できませんでした\n")
		return c.JSON(http.StatusBadRequest, "ユーザーを取得できませんでした")
	}
	var returnUsers []domain.User
	for _, user := range users {
		for _, post := range posts {
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
