package infra

import (
	"context"
	"fmt"
	"hello/server/domain"
	"hello/server/interfaces/database"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}
type Session struct {
	DryRun                   bool
	PrepareStmt              bool
	NewDB                    bool
	SkipHooks                bool
	SkipDefaultTransaction   bool
	DisableNestedTransaction bool
	AllowGlobalUpdate        bool
	FullSaveAssociations     bool
	QueryFields              bool
	Context                  context.Context
	Logger                   logger.Interface
	NowFunc                  func() time.Time
	CreateBatchSize          int
}

func DBInit() database.SqlHandler {
	env := getEnv() //envに環境変数を代入
	dsn := env.userName + ":" + env.password + "@tcp(" + env.host + ")/" + env.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("%v\n\n", dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm.Openでdbと接続している
	if err != nil {
		return nil
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = gormDB
	fmt.Println(sqlHandler.Conn) //コンソールに出力
	// migrate
	sqlHandler.Conn.AutoMigrate(&domain.Tag{}, domain.Good{}, &domain.ExternaiURL{})
	sqlHandler.Conn.AutoMigrate(&domain.Post{})
	sqlHandler.Conn.AutoMigrate(&domain.Profile{})
	sqlHandler.Conn.AutoMigrate(&domain.User{})
	return sqlHandler
}

type env struct {
	userName string
	password string
	host     string
	dbName   string
}

func getEnv() env {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envファイルの読み込みが失敗しました\n")
	}
	e := env{
		userName: os.Getenv("USERNAME"),
		password: os.Getenv("PASSWORD"),
		host:     os.Getenv("HOST"),
		dbName:   os.Getenv("DBNAME"),
	}
	return e
}
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}
func (handler *SqlHandler) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return handler.Conn.Unscoped().Delete(value, conds...)
}
func (handler *SqlHandler) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return handler.Conn.First(dest, conds...)
}
func (handler *SqlHandler) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return handler.Conn.Find(dest, conds...)
}
func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
func (handler *SqlHandler) Select(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Select(query, args...)
}
func (handler *SqlHandler) Update(column string, value interface{}) *gorm.DB {
	return handler.Conn.Update(column, value)
}
func (handler *SqlHandler) Updates(values interface{}) *gorm.DB {
	return handler.Conn.Updates(values)
}
func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}
func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}
