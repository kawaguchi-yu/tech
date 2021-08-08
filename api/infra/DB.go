package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/server/domain"
	"hello/server/interfaces/database"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	e := env{
		userName: "root",
		password: "hoge",
		host:     "db:3306",
		dbName:   "db",
	}
	return e
}
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}
func (handler *SqlHandler) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return handler.Conn.Delete(value, conds...)
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
func (handler *SqlHandler) Session(config *gorm.Session) *gorm.DB {
	return handler.Conn.Session(config)
}
func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}
func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func SetIcon(c echo.Context, db *gorm.DB) error {
	icon, err := c.FormFile("file") //cからファイルを取り出し
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

	//ここまでが画像をローカルフォルダに保存する行程、ここからがuserのiconに画像データを入れる
	email, err := ReadCookieReturnEMail(c)
	if err != nil {
		fmt.Printf("クッキー読み取りに失敗しました\n")
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := new(domain.User)
	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	db.Model(&user).Update("icon", dst.Name())
	fmt.Printf("ユーザーネーム=%v\n", user.Name)
	fmt.Printf("正常に終了しました\n" + dst.Name())
	return c.File(user.Icon)
	//例c.File(test.jpg)→test.jpgのファイルが送られる。
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
