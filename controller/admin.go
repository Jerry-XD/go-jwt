package controller

import (
	"log"
	"net/http"

	"go-jwt/forms"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a *AdminController) Create() {

}

func (a *AdminController) Read(c *gin.Context) {

	data := modelAdmin.Read()
	for _, v := range data {
		log.Println(v)
	}

	responseMsg := "Get data success !"

	c.JSON(http.StatusOK, &forms.ReadResponse{
		Message: &responseMsg,
		Data:    data,
	})
}

func (a *AdminController) Update() {

}

func (a *AdminController) Delete() {

}
