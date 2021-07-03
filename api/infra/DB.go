package infra

import (
	"fmt"
	"hello/server/domain"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBInit() error {
	env := getEnv()
	dsn := env.userName + ":" + env.password + "@tcp(" + env.host + ")/" + env.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("%v\n\n", dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
