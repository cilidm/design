package factory

type UserCreateFunc func(id int, name string) interface{}

type User struct {
	ID   int
	Name string
}

func NewUser() UserCreateFunc {
	return func(id int, name string) interface{} {
		return &User{ID: id, Name: name}
	}
}

type AdminUser struct {
	ID   int
	Name string
	Role string
}

func NewAdminUser() UserCreateFunc {
	return func(id int, name string) interface{} {
		return &AdminUser{ID: id, Name: name,Role: "admin"}
	}
}
