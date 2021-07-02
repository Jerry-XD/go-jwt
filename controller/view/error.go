package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func makeErrorResp(c *gin.Context, code int, status string, error string) {
	c.JSON(code, ErrorResp{
		Code:   code,
		Status: status,
		Error:  error,
	})
}

func MakeLoginErrorResp(c *gin.Context, error error) {
	makeErrorResp(c, http.StatusUnprocessableEntity, "Login Failed !", error.Error())
}

func MakeErrorResp(c *gin.Context, error error) {
	makeErrorResp(c, http.StatusUnprocessableEntity, "Error !", error.Error())
}
