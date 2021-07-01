package user

import "go-jwt/class"

var serviceUser = class.NewServiceUser()

func CreateUser() error {
	userID := "123456"
	err := serviceUser.Create(&userID)
	return err
}

func ReadUser() []*class.User {
	data := serviceUser.Read()
	return data
}
