package infra

import (
	"encoding/json"
	"fmt"
	"hello/server/domain"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBInit() error {
	env := getEnv() //envに環境変数を代入
	dsn := env.userName + ":" + env.password + "@tcp(" + env.host + ")/" + env.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("%v\n\n", dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm.Openでdbと接続している
	if err != nil {
		return echo.ErrInternalServerError
	}
	db = gormDB
	fmt.Println(db) //コンソールに出力
	// migrate
	db.AutoMigrate(&domain.Tag{}, domain.Good{}, &domain.ExternaiURL{})
	db.AutoMigrate(&domain.Post{})
	db.AutoMigrate(&domain.Profile{})
	db.AutoMigrate(&domain.User{})
	return nil
}
func GetDB() *gorm.DB {
	return db
}

type env struct {
	userName string
	password string
	host     string
	dbName   string
}

func getEnv() env {
	e := env{
		userName: "root",
		password: "hoge",
		host:     "db:3306",
		dbName:   "db",
	}
	return e
}
func DBCreateUser(c echo.Context, db *gorm.DB) error { //渡された値をDBに入れる
	u := new(domain.User)
	c.Bind(u) //cの中のユーザー情報をuに入れる
	rawPassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, 4)
	if err != nil {
		return echo.ErrBadRequest
	}
	u.Password = string(hashedPassword)
	fmt.Printf("%+v\n\n", u)
	result := db.Create(&u)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "メールアドレスが重複しています")
	}
	return c.JSON(http.StatusOK, "name:"+u.Name+", email:"+u.EMail+", password:"+u.Password)
}

func Login(c echo.Context, db *gorm.DB) error { //emailとpasswordでjwt入りcookie貰える
	u := new(domain.User)
	c.Bind(u)
	//ここにメルアドがdbにあるかをチェックする処理を書く
	dbPassword, err := getUser(u.EMail, db)
	if err != nil { //errの中身がnil以外なら終わる
		return c.JSON(http.StatusBadRequest, "メールアドレスが存在しませんでした")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword.Password), []byte(u.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, "パスワードが違います")
	}
	JWTToken, err := CreateJWT(u.EMail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "JWTが生成できませんでした")
	}
	cookie := CreateCookie(JWTToken)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, cookie)

}

func UserVerify(c echo.Context, db *gorm.DB) error {
	email, err := ReadCookie(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "クッキー読み取りに失敗しました")
	}
	var user domain.User
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "メールアドレスが存在しません")
	}
	return c.JSON(http.StatusOK, user)
}
func getUser(email string, db *gorm.DB) (domain.User, error) {
	var user domain.User
	if err := db.First(&user, "e_mail = ?", email).Error; err != nil {
		fmt.Printf("メールアドレスが存在しませんでした\n")
		return user, echo.ErrBadRequest
	}
	fmt.Printf("ユーザー%vユーザー\n", user)
	return user, nil
}
func GetUserModel(b io.ReadCloser) (domain.User, error) {
	var jsonData = make(map[string]string) //空っぽのmapを作る
	var user domain.User
	//デコードしてio.Reader型に変換する
	if err := json.NewDecoder(b).Decode(&jsonData); err != nil {
		return user, echo.ErrBadRequest
	}
	if jsonData == nil {
		return user, echo.ErrInternalServerError
	}

	name := jsonData["Name"]
	eMail := jsonData["EMail"]
	rawPassword := []byte(jsonData["Password"])
	//bcryptでハッシュ化したパスワードをhashedPasswordに入れる
	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, 4)
	if err != nil {
		return user, echo.ErrBadRequest
	}

	user.Name = name
	user.EMail = eMail
	user.Password = string(hashedPassword)
	return user, nil
}
