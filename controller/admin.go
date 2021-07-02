package controller

import (
	"go-jwt/controller/view"
	"go-jwt/forms"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a *AdminController) Create(c *gin.Context) {
	input := &forms.AdminCreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrorResp(c, err)
		return
	}

	res, err := modelAdmin.Create(input)
	if err != nil {
		view.MakeErrorResp(c, err)
		return
	}

	view.MakeCreateSuccessResp(c, res)
}

func (a *AdminController) Read(c *gin.Context) {
	data := modelAdmin.Read()
	view.MakeReadSuccessResp(c, data)
}

func (a *AdminController) List(c *gin.Context) {
	data := modelAdmin.List()
	view.MakeReadSuccessResp(c, data)
}

func (a *AdminController) Update(c *gin.Context) {

}

func (a *AdminController) Delete(c *gin.Context) {

}

func (a *AdminController) Login(c *gin.Context) {
	inp := &forms.LoginInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeLoginErrorResp(c, err)
		return
	}
	accessToken, err := jwt.GetAdminAccessToken(inp)
	refreshToken, err := jwt.GetAdminRefreshToken(inp)
	if err != nil {
		view.MakeLoginErrorResp(c, err)
		return
	}

	view.MakeLoginSuccessResp(c, &forms.LoginResponse{
		Message:      "Login Success !",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
