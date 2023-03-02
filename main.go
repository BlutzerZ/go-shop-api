package main

import (
	"go-shop-api/configs"
	"go-shop-api/middleware"
	"go-shop-api/routes"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()

	authMiddleware, err := middleware.JWTMiddleware()
	if err != nil {
		panic("Failed to initialize auth middleware")
	}

	r.POST("/test/login", authMiddleware.LoginHandler)
	r.GET("/hello/world", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		username := claims["Username"].(string)
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"text":     "yooo hello world from" + username,
		})
	})

	// route
	routes.UserRoute(r)
	routes.ProductRoutes(r)

	r.Run()
}
