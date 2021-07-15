package infra

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateJWT(EMail string) (string, error) { //認証が通ればCreateJWTで正規ユーザーであることを保証する
	token := jwt.New(jwt.GetSigningMethod("HS256")) //headerのセット
	token.Claims = jwt.MapClaims{
		"email": EMail,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定
	}
	tokenString, err := token.SignedString([]byte("kirikiri")) //電子署名
	if err != nil {
		return tokenString, echo.ErrBadRequest
	}
	return tokenString, nil
}

// func VerifyToken(tokenString string) (*jwt.Token, error) {
//     // jwtの検証

// }
func CreateCookie(JWTToken string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = JWTToken
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true

	return cookie
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cookie.Name)
	return c.String(http.StatusOK, "read a cookie:\n"+cookie.Value)
}
