package middlewares

import (
	"errors"
	"net/http"

	"go-jwt/config"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var conf = *config.New()

type Middleware struct{}

func (m *Middleware) AccessTokenAdmin() gin.HandlerFunc {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:              []byte(conf.App.SigningKeyAdminAccess),
		SigningAlgorithm: "HS256",
	})
	return authMiddleware.MiddlewareFunc()
}

func (m *Middleware) CheckAccessTokenAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Request.Header.Get("Authorization")
		//log.Println(token)

		res := m.GetClaimToken(c, "role")
		//log.Println("role :", res)
		if res != "admin" {
			_ = c.AbortWithError(http.StatusUnauthorized, errors.New(http.StatusText(http.StatusOK)))
		}
	}
}

func (m *Middleware) GetClaimToken(c *gin.Context, claimName string) (res string) {
	claims := jwt.ExtractClaims(c)
	if claims[claimName] != nil {
		res = claims[claimName].(string)
	}
	return res
}
