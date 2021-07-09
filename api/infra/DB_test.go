package infra

import (
	"encoding/json"
	"hello/server/testdata"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	e := echo.New()                      //Echoインスタンスを作成し代入
	body := GetCreateUserPost()          //tanaka情報をbodyに代入
	userJSON, err := json.Marshal(&body) //marshalを使ってtanaka情報を構造体からjsonに変換する
	if err != nil {
		t.Fail()
	}
	//echoを使ったhttpで送ったり返したりの処理
	//疑似的なhttpリクエストを受けることができる、PostされるuserJsonに対してstrings.NewReaderで読み込む
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) //ヘッダーの設定
	rec := httptest.NewRecorder()                                    //レスポンスを受け止める？
	c := e.NewContext(req, rec)                                      //
	assert.NotNil(t, c.Request().Body)                               //指定されたオブジェクトがnilかどうかを判別する

	db, mock, err := testdata.GetMockDB()
	mock.ExpectExec(regexp.QuoteMeta(`CREATE`)).WillReturnResult(sqlmock.NewResult(1, 10)) //第一引数は主キーのIDの値
	mock.ExpectExec(regexp.QuoteMeta(`INSERT`)).WillReturnResult(sqlmock.NewResult(1, 10)) //第二引数は本クエリのよって影響を受けるカラムの数

	if err != nil {
		t.Log(err)
		t.Log("mockdb down")
		t.Fail()
	}
	if err := testdata.TestMigrate(db, mock); err != nil {
		t.Log(err)
		t.Log("migrate down")
		t.Fail()
	}
	if err := DBCreateUser(c, db); err != nil {
		t.Log(err)
		t.Log("makeuser down")
		t.Fail()
	}
}

type CreateUserData struct {
	Name     string
	EMail    string
	Password string
}

//CreateUserDataという構造体のテスト用データを返す関数
func GetCreateUserPost() CreateUserData {
	body := CreateUserData{
		Name:     "tanaka",
		EMail:    "tanaka@gmail.com",
		Password: "tanaka1111",
	}
	return body
}
func TestGetUserModel(t *testing.T) {
	body := GetCreateUserPost()              //tanaka情報をbodyに入れる
	jsonBody, err := json.Marshal(&body)     //jsonへと変換する
	assert.Nil(t, err)                       //nilかどうかをテスト
	r := strings.NewReader(string(jsonBody)) //stringのt情報を
	readCloser := ioutil.NopCloser(r)        //io.ReadCloserインターフェースCloseを満たすオブジェクトを返す
	_, errUser := GetUserModel(readCloser)   //
	assert.Nil(t, errUser)                   //nilかどうかをテスト
}
