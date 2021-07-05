package infra

import (
	"fmt"
	"hello/server/domain"
	"net/http"

	"github.com/labstack/echo/v4"
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

func CreateUser(c echo.Context) error {
	u := new(domain.User)
	c.Bind(u)
	fmt.Printf("%+v\n\n", u)
	DBCreateData(*u)
	return c.JSON(http.StatusOK, "name:"+u.Name+", email:"+u.EMail+", password:"+u.Password)
}

func DBCreateData(u domain.User) error {
	db := GetDB()
	user := u
	db.Create(&user)
	db.Find(&user, "ID = ?", 1)
	fmt.Println("userの値は", user)
	return nil
}
