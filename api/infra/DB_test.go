package infra

import (
	"database/sql/driver"
	"hello/server/domain"
	"hello/server/testdata"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// func newDummyHandler(db *gorm.DB) database.SqlHandler {
// 	sqlHandler := new(SqlHandler)
// 	sqlHandler.Conn = db
// 	return sqlHandler
// }
func TestCreate(t *testing.T) {
	body := GetCreateUser() //tanaka情報をbodyに代入
	db, mock, err := testdata.GetMockDB()
	r := SqlHandler{Conn: db}
	query := "INSERT"
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 10)) //第一引数は主キーのIDの値

	if err != nil {
		t.Log(err)
		t.Log("mockdb down")
		t.Fail()
	}
	if err := r.Create(body).Error; err != nil {
		t.Log(err)
		t.Log("makeuser down")
		t.Fail()
	}
}
func TestDelete(t *testing.T) {
	body := GetCreateUser() //tanaka情報をbodyに代入
	db, mock, err := testdata.GetMockDB()
	r := SqlHandler{Conn: db}
	query := "DELETE FROM `users` WHERE name = ?"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(body.Name).
		WillReturnResult(sqlmock.NewResult(1, 2))
	if err != nil {
		t.Log(err)
		t.Log("mockdb down")
		t.Fail()
	}
	//gormのdeleteが正しく実行されているかをテスト
	if err := r.Delete(&domain.User{}, "name = ?", body.Name).Error; err != nil {
		t.Log(err) //正常でなければエラーが出てここで弾く
		t.Log("delete down")
		t.Fail()
	}
}
func TestFirst(t *testing.T) {
	body := GetCreateUser() //tanaka情報をbodyに代入
	user := domain.User{}
	db, mock, err := testdata.GetMockDB()
	r := SqlHandler{Conn: db}
	query := "SELECT * FROM `users` WHERE name = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)). //mockdbは次に来る(r.First())クエリがこれであることを期待する
							WithArgs(body.Name).                                        //クエリに?があればそこに入る値を指定する
							WillReturnRows(sqlmock.NewRows([]string{"name", "e_mail"}). //これが返ってくることを期待する
															AddRow(body.Name, body.EMail)) //mockdbにselectで貰ってくるデータを入れる
	name := "tanaka"
	email := "tanaka@gmail.com"
	if err != nil {
		t.Log(err)
		t.Log("mockdb down")
		t.Fail()
	}
	if err := r.First(&user, "name = ?", body.Name).Error; err != nil {
		t.Log(err)
		t.Log("get down")
		t.Fail()
	}
	if user.Name != name || user.EMail != email { //dbから貰った値が一致しているか(変わってないか、変な値を貰ってきてないか)を判定する
		t.Log("取得結果不一致")
		t.Fail()
	}
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func TestUpdate(t *testing.T) {
	user := GetCreateUser() //tanaka情報をbodyに代入
	name := "tarou"
	user.ID = 1 //gorm.dbの番号
	db, mock, err := testdata.GetMockDB()
	r := SqlHandler{Conn: db}
	query := "UPDATE `users` SET `name`=?,`updated_at`=? WHERE `id` = ?"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(name, AnyTime{}, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 10))
	if err != nil {
		t.Log(err)
		t.Log("mockdb down")
		t.Fail()
	}
	//gormのdeleteが正しく実行されているかをテスト

	if err := r.Model(&user).Update("name", name).Error; err != nil {
		t.Log(err) //正常でなければエラーが出てここで弾く
		t.Log("update down")
		t.Fail()
	}
	if user.Name != "tarou" {
		t.Log("取得結果不一致")
		t.Fail()
	}
}

type CreateUserData struct {
	Name  string
	EMail string
}

//CreateUserDataという構造体のテスト用データを返す関数
func GetCreateUser() *domain.User {
	body := &domain.User{
		Name:  "tanaka",
		EMail: "tanaka@gmail.com",
	}
	return body
}

// func TestGetUserModel(t *testing.T) {
// 	body := GetCreateUser()                  //tanaka情報をbodyに入れる
// 	jsonBody, err := json.Marshal(&body)     //jsonへと変換する
// 	assert.Nil(t, err)                       //nilかどうかをテスト
// 	r := strings.NewReader(string(jsonBody)) //stringのt情報を
// 	readCloser := ioutil.NopCloser(r)        //io.ReadCloserインターフェースCloseを満たすオブジェクトを返す
// 	_, errUser := GetUserModel(readCloser)   //
// 	assert.Nil(t, errUser)                   //nilかどうかをテスト
// }
