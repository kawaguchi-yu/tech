package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
)

func GuestUserCheck(userName string) error {
	const GuestUserName = "Guest User"
	if userName == GuestUserName {
		return errors.New("GuestUserはユーザーを削除する権限がありません")
	}
	return nil
}
func GetJWTToken() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envファイルの読み込みが失敗しました\n")
		return "error", err
	}
	encryptionKey := os.Getenv("ENCRYPTIONKEY")
	return encryptionKey, nil
}
func ReadCookieReturnEMail(c Context) (string, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Printf("クッキーを読み込めませんでした%v\n", cookie)
		return "error", err
	}
	encryptionKey, err := GetJWTToken()
	if err != nil {
		return "error", err
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(encryptionKey), nil
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
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envファイルの読み込みが失敗しました\n")
	}
	awsAccesskey := os.Getenv("AWSACCESSKEY")
	awsSecretkey := os.Getenv("AWSSECRETKEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsAccesskey, awsSecretkey, ""),
		Region:      aws.String("ap-northeast-1"),
	}))
	fmt.Printf("111usericon=%v\n", userIcon)
	svc := s3.New(sess)
	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("techer-s3-001"),
		Key:    aws.String(userIcon),
	})
	if err != nil {
		fmt.Printf("awsからデータを貰えませんでした error=%v\n", err)
		return "error"
	}
	obj := object.Body
	defer obj.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, obj)
	object.Body.Read(buf.Bytes())
	userIcon = base64.StdEncoding.EncodeToString(buf.Bytes())
	return userIcon
}
