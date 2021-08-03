package infra

import (
	"fmt"

	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateJWT(EMail string) (string, error) { //認証が通ればCreateJWTで正規ユーザーであることを保証する

	claims := jwt.StandardClaims{
		Issuer:    EMail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("kirikiri")) //電子署名
	if err != nil {
		return tokenString, echo.ErrBadRequest
	}
	return tokenString, nil
}

func CreateCookie(JWTToken string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = JWTToken
	cookie.Expires = time.Now().Add(10 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true

	return cookie
}

type Claims struct {
	jwt.StandardClaims
}

func ReadCookieReturnEMail(c echo.Context) (string, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Printf("クッキーを読み込めませんでした%v\n", cookie)
		return "error", echo.ErrBadRequest
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("kirikiri"), nil
	})
	if err != nil || !token.Valid {
		fmt.Printf("パルスに失敗しました\n")
		return cookie.Value, echo.ErrBadRequest
	}
	claims := token.Claims.(*Claims)
	email := claims.Issuer
	return email, nil
}
