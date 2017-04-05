package Models

import (
	"net/http"

	"github.com/mholt/binding"
)

type User struct {
	UserName  string
	Email     string
	Password  string
	Message   string
	Remember  bool
	IsLogedIn bool
}

func (user *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&user.Email:    "email",
		&user.UserName: "username",
		&user.Password: "password",
		&user.Message: binding.Field{
			Form:     "message",
			Required: true,
		},
	}
}
