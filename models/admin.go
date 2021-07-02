package models

import (
	"errors"

	"go-jwt/forms"
	"go-jwt/interfaces"
)

func NewServiceAdmin() (service interfaces.ServiceAdmin) {
	return &crud{}
}

type crud struct{}

func (crud *crud) Create(data *string) (err error) {
	err = errors.New("Error")
	return err
}

func (crud *crud) Read() (admin []*forms.Admin) {
	return []*forms.Admin{
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
