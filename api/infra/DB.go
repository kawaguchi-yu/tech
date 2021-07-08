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

func CreateUser(c echo.Context) error { //中身をDBに入れてJSONで返す
	u := new(domain.User) //uにUser構造体を代入
	c.Bind(u)             //c(引数)をuの中に入れる
	fmt.Printf("%+v\n\n", u)
	return c.JSON(http.StatusOK, "name:"+u.Name+", email:"+u.EMail+", password:"+u.Password)
	//uの中に何が入っているかをJSON形式で返す
}

func DBCreateUser(c echo.Context, db *gorm.DB) error { //渡された値をDBに入れる
	u := new(domain.User)
	c.Bind(u)
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
