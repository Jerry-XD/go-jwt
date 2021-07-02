package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResp struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func makeSuccessResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, SuccessResp{
		Code:   code,
		Status: msg,
		Data:   data,
	})
}

func MakeLoginSuccessResp(c *gin.Context, data interface{}) {
	makeSuccessResp(c, http.StatusOK, "Login Success !", data)
}

func MakeReadSuccessResp(c *gin.Context, data interface{}) {
	makeSuccessResp(c, http.StatusOK, "Get Data Success !", data)
}

func MakeCreateSuccessResp(c *gin.Context, data interface{}) {
	makeSuccessResp(c, http.StatusOK, "Create Data Success !", data)
}
