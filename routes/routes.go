package routes

import (
	"fmt"

	"go-jwt/config"
	"go-jwt/controller"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct {
	admin controller.AdminController
}

var (
	r                          Route
	conf                       = *config.New()
	middleware                 = new(middlewares.Middleware)
	middlewareAccessTokenAdmin gin.HandlerFunc
	router                     *gin.Engine
)

func NewRouter() *gin.Engine {
	return router
}

func init() {
	router = gin.Default()

	// middleware
	middlewareAccessTokenAdmin = middleware.AccessTokenAdmin()

	// Status check
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK !",
		})
	})

	// Router Group
	groupAdmin := router.Group("/admin")

	// Authentication
	groupAdmin.POST("/login", r.admin.Login)

	// Route admin
	admin := groupAdmin.Group("", middlewareAccessTokenAdmin, middleware.CheckAccessTokenAdmin())
	{
		admin.POST("/data", r.admin.Create)
		admin.GET("/data/:id", r.admin.Read)
		admin.GET("/data", r.admin.List)
		admin.PUT("/data", r.admin.Update)
		admin.DELETE("/data", r.admin.Delete)
	}

}

func Start() {
	err := router.Run(fmt.Sprintf(":%s", conf.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
