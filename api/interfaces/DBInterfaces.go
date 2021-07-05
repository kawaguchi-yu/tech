package interfaces

import (
	"fmt"
	"hello/server/domain"
	"hello/server/infra"
)

func DBCreateData() error {
	db := infra.GetDB()
	user := domain.User{Name: "kiri", EMail: "kiri@gmail.com", Password: "aaaa"}
	db.Create(&user)
	db.Find(&user, "ID = ?", 1)
	fmt.Println("userの値は", user)
	return nil
}
