package Models

import (
	"log"
	"net/http"

	"github.com/codeNut247/golang_website/DataBinding"
)

type Authorization struct {
	UserName string
	Email    string
	Password string
	Remember string
}

func (am *Authorization) fillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := binding.SetField(am, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func BindModel(req *http.Request) *Authorization {
	formData := make(map[string]interface{})
	formData["UserName"] = req.FormValue("username")
	formData["Email"] = req.FormValue("email")
	formData["Password"] = req.FormValue("password")
	formData["Remember"] = req.FormValue("remember")

	result := &Authorization{}
	err := result.fillStruct(formData)
	if err != nil {
		log.Fatalln("Data could not be bound to Struct")
	}

	return result
}

/*
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
*/
