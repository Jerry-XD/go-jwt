package routes

import (
	"fmt"
	"log"

	"go-jwt/config"
	"go-jwt/controller"
	"go-jwt/forms"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct {
	admin controller.AdminController
}

var (
	r                         Route
	conf                      = *config.New()
	middleware                = new(middlewares.Middleware)
	middlewareAccesTokenAdmin gin.HandlerFunc
	router                    *gin.Engine
)

func init() {
	router = gin.Default()

	// middleware
	middlewareAccesTokenAdmin = middleware.AccessTokenAdmin()

	// Status check
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK !",
		})
	})

	// Router Group
	groupAdmin := router.Group("/admin")

	// Authentication
	groupAdmin.POST("/login", func(c *gin.Context) {
		inp := &forms.LoginInput{}
		if err := c.ShouldBindJSON(inp); err != nil {
			log.Println(err)
			return
		}
		accessToken, err := middleware.GetAdminAccessToken(inp)
		refreshToken, err := middleware.GetAdminRefreshToken(inp)
		log.Println(accessToken)
		log.Println(refreshToken)
		log.Println(err)

		c.JSON(200, gin.H{
			"message":       "Login Success !",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	})

	// Route admin
	admin := groupAdmin.Group("", middlewareAccesTokenAdmin, middleware.CheckAccessTokenAdmin())
	{
		admin.GET("/data", r.admin.Read)
	}

}

func Start() {
	_ = router.Run(fmt.Sprintf(":%s", conf.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
