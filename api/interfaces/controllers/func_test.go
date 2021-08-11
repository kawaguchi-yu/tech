package controllers

// import (
// 	"encoding/json"
// 	"hello/server/domain"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestReadCookieReturnEMail(t *testing.T) {
// 	e := echo.New()
// 	body := getTestUser()
// 	userJSON, err := json.Marshal(&body)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(userJSON)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	ReadCookieReturnEMail(c)

// }
// func getTestUser() domain.User {
// 	user := domain.User{
// 		Name:     "tanaka",
// 		EMail:    "tanaka@gmail.com",
// 		Password: "tanaka1234",
// 	}
// 	return user
// }
