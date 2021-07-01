package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go-jwt/user"

	jwt "github.com/appleboy/gin-jwt/v2"
	jwts "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReadResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	SigningKeyAdminAccess  = "some-salt-access-key"
	SigningKeyAdminRefresh = "some-salt-refresh-key"
)

var (
	middlewareAccestoken gin.HandlerFunc
)

func main() {
	r := gin.Default()

	// middleware
	middlewareAccestoken = AccessToken()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/data", middlewareAccestoken, CheckAccessTokenSuper(), func(c *gin.Context) {
		data := user.ReadUser()
		responseMsg := "Get data success !"
		resp := &ReadResponse{
			Message: &responseMsg,
			Data:    data,
		}
		c.JSON(200, resp)
	})

	r.POST("/login", func(c *gin.Context) {
		inp := &LoginInput{}
		if err := c.ShouldBindJSON(inp); err != nil {
			log.Println(err)
			return
		}
		accessToken, err := GetAdminAccessToken(inp)
		refreshToken, err := GetAdminRefreshToken(inp)
		log.Println(accessToken)
		log.Println(refreshToken)
		log.Println(err)

		c.JSON(200, gin.H{
			"message":       "Login Success !",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	})

	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetAdminAccessToken(data *LoginInput) (string, error) {
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

func GetAdminRefreshToken(data *LoginInput) (string, error) {
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

func AccessToken() gin.HandlerFunc {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:              []byte(SigningKeyAdminAccess),
		SigningAlgorithm: "HS256",
	})
	return authMiddleware.MiddlewareFunc()
}

func CheckAccessTokenSuper() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		log.Println(token)
		//test := jwt.GetToken(c)
		test := "role"
		role := GetClaimToken(c, test)
		//log.Println("role :", role)
		//log.Println("Token :", token)

		if role != "admin" {
			_ = c.AbortWithError(200, errors.New(http.StatusText(http.StatusOK)))
		}
	}
}

func GetClaimToken(c *gin.Context, claimName string) (res string) {
	claims := jwt.ExtractClaims(c)
	if claims[claimName] != nil {
		res = claims[claimName].(string)
	}
	return res
}
