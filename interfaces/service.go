package interfaces

import "go-jwt/forms"

type ServiceAdmin interface {
	Create(data *string) (err error)
	Read() (admin []*forms.Admin)
	Update(data string) (result string)
	Delete(data string) (err error)
}
