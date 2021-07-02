package middlewares

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go-jwt/forms"

	jwt "github.com/appleboy/gin-jwt/v2"
	jwts "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	SigningKeyAdminAccess  = "some-salt-access-key"
	SigningKeyAdminRefresh = "some-salt-refresh-key"
)

type Middleware struct{}

func (m *Middleware) GetAdminAccessToken(data *forms.LoginInput) (string, error) {
	jwtToken := jwts.NewWithClaims(jwts.SigningMethodHS256, jwts.MapClaims{
		"username": data.Username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Minute * time.Duration(2400)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(SigningKeyAdminAccess))
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func (m *Middleware) GetAdminRefreshToken(data *forms.LoginInput) (string, error) {
	jwtToken := jwts.NewWithClaims(jwts.SigningMethodHS256, jwts.MapClaims{
		"username": data.Username,
		"exp":      time.Now().Add(time.Hour * time.Duration(360)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(SigningKeyAdminRefresh))
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func (m *Middleware) AccessTokenAdmin() gin.HandlerFunc {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:              []byte(SigningKeyAdminAccess),
		SigningAlgorithm: "HS256",
	})
	return authMiddleware.MiddlewareFunc()
}

func (m *Middleware) CheckAccessTokenAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		log.Println(token)
		//test := jwt.GetToken(c)
		test := "role"
		role := m.GetClaimToken(c, test)
		//log.Println("role :", role)
		//log.Println("Token :", token)

		if role != "admin" {
			_ = c.AbortWithError(200, errors.New(http.StatusText(http.StatusOK)))
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
