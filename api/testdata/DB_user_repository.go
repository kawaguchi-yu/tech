package testdata

import (
	"hello/server/domain"
)

type UserRepository interface {
	CreateUser(domain.User) error
}
