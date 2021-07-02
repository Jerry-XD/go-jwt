package models

import (
	"errors"
	"log"

	"go-jwt/forms"
	"go-jwt/interfaces"
)

func NewServiceAdmin() (service interfaces.ServiceAdmin) {
	return &crud{}
}

type crud struct{}

func (crud *crud) Create(input *forms.AdminCreateInput) (res *forms.AdminCreateInput, err error) {
	// Create Data to Database
	log.Println("Create Success !")
	return input, nil
}

func (crud *crud) Read() (admin *forms.Admin) {
	return admin
}

func (crud *crud) List() (admin []*forms.Admin) {
	// Query data from database
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
