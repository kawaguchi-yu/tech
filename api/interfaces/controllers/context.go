package controllers

import (
	"mime/multipart"
	"net/http"
)

type Context interface {
	SetCookie(cookie *http.Cookie)
	Cookie(name string) (*http.Cookie, error)
	FormFile(name string) (*multipart.FileHeader, error)
	File(file string) error
	Bind(interface{}) error
	JSON(int, interface{}) error
}
