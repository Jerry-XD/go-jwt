package util

import (
	"time"

	"go-jwt/config"
	"go-jwt/forms"

	jwts "github.com/dgrijalva/jwt-go"
)

var conf = *config.New()

type JWT interface {
	GetAdminAccessToken(data *forms.LoginInput) (string, error)
	GetAdminRefreshToken(data *forms.LoginInput) (string, error)
}

type jwt struct{}

func NewJWT() JWT {
	return &jwt{}
}

func (j *jwt) GetAdminAccessToken(data *forms.LoginInput) (string, error) {
	jwtToken := jwts.NewWithClaims(jwts.SigningMethodHS256, jwts.MapClaims{
		"username": data.Username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Minute * time.Duration(conf.App.AccessTokenExp)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(conf.App.SigningKeyAdminAccess))
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func (j *jwt) GetAdminRefreshToken(data *forms.LoginInput) (string, error) {
	jwtToken := jwts.NewWithClaims(jwts.SigningMethodHS256, jwts.MapClaims{
		"username": data.Username,
		"exp":      time.Now().Add(time.Minute * time.Duration(conf.App.RefreshTokenExp)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(conf.App.SigningKeyAdminRefresh))
	if err != nil {
		return "", err
	}
	return strToken, nil
}
