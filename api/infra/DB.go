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

// func Login(c echo.Context, db *gorm.DB) error { //emailとpasswordでjwt入りcookie貰える
// 	u := new(domain.User)
// 	c.Bind(u)
// 	//ここにメルアドがdbにあるかをチェックする処理を書く
// 	dbPassword, err := getUser(u.EMail, db)
// 	if err != nil { //errの中身がnil以外なら終わる
// 		return c.JSON(http.StatusBadRequest, "メールアドレスが存在しませんでした")
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword.Password), []byte(u.Password)); err != nil {
// 		return c.JSON(http.StatusBadRequest, "パスワードが違います")
// 	}
// 	JWTToken, err := CreateJWT(u.EMail)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "JWTが生成できませんでした")
// 	}
// 	cookie := CreateCookie(JWTToken)
// 	c.SetCookie(cookie)
// 	return c.JSON(http.StatusOK, cookie)
// }
// func DeleteUser(c echo.Context, db *gorm.DB) error {
// 	email, err := ReadCookieReturnEMail(c) //クッキーからemailを読み取る
// 	if err != nil {
// 		fmt.Printf("ユーザーの削除権限がありません\n")
// 		return c.JSON(http.StatusBadRequest, "ユーザーの削除権限がありません")
// 	}
// 	user := new(domain.User)
// 	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, "db接続に失敗しました\n")
// 	} //dbから現在のユーザーデータを取得する
// 	db.Delete(&domain.User{}, user.ID)
// 	posts := new(domain.Post)
// 	db.Where("user_id = ?", user.ID).Find(&posts)
// 	fmt.Printf("%v\n", posts)
// 	db.Delete(&domain.Post{}, posts.ID)
// 	fmt.Printf("userから%vIDを削除しました\n", user.ID)
// 	return c.JSON(http.StatusOK, "ユーザーを削除しました。")
// }
// func Logout(c echo.Context, db *gorm.DB) error {
// 	cookie, err := c.Cookie("jwt")
// 	if err != nil {
// 		fmt.Printf("クッキー読み取りに失敗しました\n")
// 		return c.JSON(http.StatusBadRequest, cookie)
// 	}
// 	cookie.Expires = time.Now()
// 	c.SetCookie(cookie)
// 	return c.JSON(http.StatusOK, cookie)
// }

// func getIcon(userIcon string) string {
// 	os.Chdir("img")
// 	file, err := os.Open(userIcon)
// 	if err != nil {
// 		fmt.Printf("データを開けませんでした\n")
// 		return "error"
// 	}
// 	defer file.Close()
// 	fi, err := file.Stat() //FileInfo interface
// 	if err != nil {
// 		fmt.Printf("データ取得に失敗しました\n")
// 		return "error"
// 	}
// 	size := fi.Size() //ファイルサイズ
// 	data := make([]byte, size)
// 	file.Read(data)
// 	userIcon = base64.StdEncoding.EncodeToString(data)
// 	return userIcon
// }

// func ReadCookieReturnUser(c echo.Context, db *gorm.DB) error {
// 	email, err := ReadCookieReturnEMail(c)
// 	if err != nil {
// 		fmt.Printf("クッキー読み取りに失敗しました\n")
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	var user domain.User
// 	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
// 		fmt.Printf("emailが存在しませんでした\n")
// 		return c.JSON(http.StatusBadRequest, "emailが存在しませんでした")
// 	}
// 	user.Icon = getIcon(user.Icon)
// 	if user.Icon == "error" {
// 		return c.JSON(http.StatusBadRequest, "base64エンコードに失敗しました")
// 	}
// 	fmt.Printf("user.Name=%v\n", user.Name)
// 	return c.JSON(http.StatusOK, user)
// }
// func ReturnAllUserPost(c echo.Context, db *gorm.DB) error { //全user情報を渡す
// 	os.Chdir("img")
// 	var users []domain.User
// 	db.Find(&users)
// 	var posts []domain.Post
// 	db.Find(&posts)
// 	var returnUsers []domain.User
// 	for _, user := range users {
// 		for _, post := range posts {
// 			if post.UserID == user.ID {
// 				user.Posts = append(user.Posts, post)
// 			}
// 		}
// 		if user.Icon != "" {
// 			fmt.Printf("userIcon=%vuserName=%v\n", user.Icon, user.Name)
// 			file, err := os.Open(user.Icon)
// 			if err != nil {
// 				return c.JSON(http.StatusBadRequest, "データを開けませんでした")
// 			}
// 			defer file.Close()
// 			fi, err := file.Stat() //FileInfo interface
// 			if err != nil {
// 				fmt.Printf("データ取得に失敗しました\n")
// 				return c.JSON(http.StatusBadRequest, "データ取得に失敗しました")
// 			}
// 			size := fi.Size() //ファイルサイズ
// 			data := make([]byte, size)
// 			file.Read(data)
// 			user.Icon = base64.StdEncoding.EncodeToString(data)
// 		}
// 		if user.Posts != nil {
// 			returnUsers = append(returnUsers, user)
// 		}
// 	}
// 	fmt.Printf("ReturnAllUserは正常に終了しました\n")
// 	return c.JSON(http.StatusOK, returnUsers)
// }
// func ReadURLReturnUserPost(c echo.Context, db *gorm.DB) error {
// 	u := new(domain.User)
// 	c.Bind(u) //これでフロントエンド側のurlの名前を受け取る
// 	var user domain.User
// 	if err := db.First(&user, "name=?", u.Name).Error; err != nil {
// 		fmt.Printf("ユーザー取得に失敗しました%v\n", u.Name)
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	var posts []domain.Post
// 	db.Where("user_id = ?", user.ID).Find(&posts)
// 	for _, post := range posts {
// 		if post.UserID == user.ID {
// 			user.Posts = append(user.Posts, post)
// 		}
// 	}
// 	fmt.Printf("userIcon=%vuserName=%v\n", user.Icon, user.Name)
// 	user.Icon = getIcon(user.Icon)
// 	if user.Icon == "" {
// 		return c.JSON(http.StatusOK, "base64エンコードに失敗しました")
// 	}
// 	fmt.Printf("userIDは%v\n", user.ID)
// 	fmt.Printf("Post引き出し処理は正常に終了しました\n")
// 	return c.JSON(http.StatusOK, user)
// }

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

// func CreatePost(c echo.Context, db *gorm.DB) error {
// 	post := new(domain.Post)
// 	c.Bind(post)
// 	email, err := ReadCookieReturnEMail(c)
// 	if err != nil {
// 		fmt.Printf("クッキー読み取りに失敗しました\n")
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	user := new(domain.User)
// 	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	post.UserID = user.ID
// 	user.Posts = []domain.Post{{
// 		UserID:       post.UserID,
// 		Title:        post.Title,
// 		Answer:       post.Answer,
// 		WrongAnswer1: post.WrongAnswer1,
// 		WrongAnswer2: post.WrongAnswer2,
// 		WrongAnswer3: post.WrongAnswer3,
// 		Explanation:  post.Explanation,
// 		Tags:         post.Tags,
// 		Goods:        post.Goods,
// 	}}
// 	db.Select("Posts").Updates(&user)
// 	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&user).Association("Posts").Append(&user.Posts)
// 	if err != nil {
// 		fmt.Printf("%v/n", err)
// 	}
// 	fmt.Printf("処理は正常に終了しました\n")
// 	return c.JSON(http.StatusOK, post)
// }
// func UpDataPost(c echo.Context, db *gorm.DB) error {
// 	post := new(domain.Post)
// 	c.Bind(post)
// 	email, err := ReadCookieReturnEMail(c)
// 	if err != nil {
// 		fmt.Printf("クッキー読み取りに失敗しました\n")
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	user := new(domain.User)
// 	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, nil)
// 	}
// 	if user.ID == post.UserID {
// 		result := db.Model(&domain.Post{}).Where("Id =?", post.ID).Updates(post)
// 		if result.Error != nil {
// 			fmt.Printf("アップデートに失敗しました\n")
// 			return c.JSON(http.StatusBadRequest, "アップデートに失敗しました")
// 		}
// 		fmt.Printf("postの%vをアップデートしました\n", post.ID)
// 	} else {
// 		fmt.Printf("編集権限がありませんuser.ID=%vpost.UserID=%v\n", user.ID, post.UserID)
// 		return c.JSON(http.StatusBadRequest, "編集権限がありません")
// 	}
// 	return c.JSON(http.StatusOK, "正常に終了しました")
// }
// func DeletePost(c echo.Context, db *gorm.DB) error {
// 	type postData struct {
// 		ID     uint
// 		UserID uint
// 	}
// 	post := new(postData)
// 	c.Bind(post)
// 	fmt.Printf("postid%v\n", post)         //postのIDとUserIDが入ってる
// 	email, err := ReadCookieReturnEMail(c) //クッキーからemailを読み取る
// 	if err != nil {
// 		fmt.Printf("記事の削除権限がありません\n")
// 		return c.JSON(http.StatusBadRequest, "削除権限がありません")
// 	}
// 	user := new(domain.User)
// 	if err := db.First(&user, "e_mail=?", email).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, "db接続に失敗しました\n")
// 	} //dbから現在のユーザーデータを取得する
// 	if user.ID == post.UserID {
// 		db.Delete(&domain.Post{}, post.ID)
// 		fmt.Printf("postから%vIDを削除しました\n", post.ID)
// 	} else {
// 		fmt.Printf("削除権限がありませんuser.ID=%vpost.UserID=%v\n", user.ID, post.UserID)
// 		return c.JSON(http.StatusBadRequest, "削除権限がありません")
// 	}
// 	fmt.Printf("正常に終了しました%v\n", post.ID)
// 	return c.JSON(http.StatusOK, "正常に終了しました")
// }
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

// func DBCreateUser(c echo.Context, db *gorm.DB) error { //渡された値をDBに入れる
// 	u := new(domain.User)
// 	c.Bind(u) //cの中のユーザー情報をuに入れる
// 	rawPassword := []byte(u.Password)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(rawPassword, 4)
// 	if err != nil {
// 		return echo.ErrBadRequest
// 	}
// 	u.Password = string(hashedPassword)
// 	fmt.Printf("%+v\n", u)
// 	u.Icon = ("dog_out.png")
// 	result := db.Create(&u)
// 	if result.Error != nil {
// 		return c.JSON(http.StatusBadRequest, "名前かメールアドレスが重複しています")
// 	}
// 	return c.JSON(http.StatusOK, "ユーザー登録完了！")
// }
