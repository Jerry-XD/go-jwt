package controller

import (
	"go-jwt/models"
	"go-jwt/util"
)

var (
	modelAdmin = models.NewServiceAdmin()
	jwt        = util.NewJWT()
)
