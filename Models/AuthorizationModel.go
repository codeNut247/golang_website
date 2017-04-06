package Models

import binding "github.com/codeNut247/golang_website/DataBinding"

type UserModel struct {
	UserName  string
	Email     string
	Password  string
	Message   string
	Remember  bool
	IsLogedIn bool
}

func (um *UserModel) BindModel(m map[string]interface{}) error {
	for k, v := range m {
		err := binding.SetField(um, k, v)
		if err != nil {
			return err
		}
	}
	return nil
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
