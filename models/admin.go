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

func (crud *crud) Create(input *forms.AdminCreateInput) (res *forms.Admin, err error) {
	// Create Data to Database
	res = &forms.Admin{
		ID:    "1",
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}
	log.Println("Create Success !")

	return res, nil
}

func (crud *crud) Read(id *string) (admin *forms.Admin, err error) {
	// Query data from database
	if *id == "1" {
		admin = &forms.Admin{
			ID:    "1",
			Name:  "Test1",
			Email: "gmail.com",
			Age:   10,
		}
	} else {
		return nil, errors.New("Data not found !")
	}
	return admin, nil
}

func (crud *crud) List() (admin []*forms.Admin) {
	// Query data from database
	return []*forms.Admin{
		{
			ID:    "1",
			Name:  "Test1",
			Email: "gmail.com",
			Age:   10,
		},
		{
			ID:    "2",
			Name:  "Test2",
			Email: "gmail.com",
			Age:   20,
		},
		{
			ID:    "3",
			Name:  "Test3",
			Email: "gmail.com",
			Age:   30,
		},
	}
}

func (crud *crud) Update(data string) (result string, err error) {
	return "", nil
}

func (crud *crud) Delete(Update string) (err error) {
	err = errors.New("Error")
	return err
}
