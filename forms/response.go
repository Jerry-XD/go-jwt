package forms

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReadResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}
