package interfaces

import "go-jwt/forms"

type ServiceAdmin interface {
	Create(input *forms.AdminCreateInput) (res *forms.Admin, err error)
	Read(id *string) (admin *forms.Admin, err error)
	List() (admin []*forms.Admin)
	Update(data string) (result string, err error)
	Delete(data string) (err error)
}
