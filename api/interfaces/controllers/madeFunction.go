package controllers

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ReadCookieReturnEMail(c Context) (string, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Printf("クッキーを読み込めませんでした%v\n", cookie)
		return "error", err
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("kirikiri"), nil
	})
	if err != nil || !token.Valid {
		fmt.Printf("パルスに失敗しました\n")
		return cookie.Value, err
	}
	claims := token.Claims.(*Claims)
	email := claims.Issuer
	return email, nil
}
func getIcon(userIcon string) string {
	os.Chdir("img")
	file, err := os.Open(userIcon)
	if err != nil {
		fmt.Printf("データを開けませんでした\n")
		return "error"
	}
	defer file.Close()
	fi, err := file.Stat() //FileInfo interface
	if err != nil {
		fmt.Printf("データ取得に失敗しました\n")
		return "error"
	}
	size := fi.Size() //ファイルサイズ
	data := make([]byte, size)
	file.Read(data)
	userIcon = base64.StdEncoding.EncodeToString(data)
	return userIcon
}
