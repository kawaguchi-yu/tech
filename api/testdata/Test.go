package testdata

import (
	"hello/server/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// skip transaction for sqlmock testing
	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestMigrate(db *gorm.DB, mock sqlmock.Sqlmock) error {

	db.AutoMigrate(&domain.User{})
	return nil
}
