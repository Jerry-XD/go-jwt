package forms

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ListResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    []Admin `json:"data"`
}
