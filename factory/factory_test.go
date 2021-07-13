package factory

import (
	"fmt"
	"testing"
)

const (
	FrontUserType = iota
	AdminUserType
)
type UserType int

func CreateUser(t UserType) UserCreateFunc{
	switch t {
	case FrontUserType:
		return NewUser()
	case AdminUserType:
		return NewAdminUser()
	default:
		return NewUser()
	}
}

func TestFactoryNew(t *testing.T){
	user := CreateUser(FrontUserType)(123,"abc").(*User)
	fmt.Println(user)
	admin := CreateUser(AdminUserType)(124,"admin").(*AdminUser)
	fmt.Println(admin)
}
