package class

import "errors"

type ServiceUser interface {
	Create(data *string) (err error)
	Read() (user []*User)
	Update(data string) (result string)
	Delete(data string) (err error)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

func NewServiceUser() (service ServiceUser) {
	return &crud{}
}

type crud struct{}

func (crud *crud) Create(data *string) (err error) {
	err = errors.New("Error")
	return err
}

func (crud *crud) Read() (user []*User) {
	return []*User{
		{
			Name:  "Test1",
			Email: "gmail.com",
			Age:   10,
		},
		{
			Name:  "Test2",
			Email: "gmail.com",
			Age:   20,
		},
		{
			Name:  "Test3",
			Email: "gmail.com",
			Age:   30,
		},
	}
}

func (crud *crud) Update(data string) (result string) {
	return ""
}

func (crud *crud) Delete(Update string) (err error) {
	err = errors.New("Error")
	return err
}
