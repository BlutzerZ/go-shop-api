package main

import (
	"go-shop-api/configs"
	"go-shop-api/middleware"
	"go-shop-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()

	authMiddleware, err := middleware.JWTMiddleware()
	if err != nil {
		panic("Failed to initialize auth middleware")
	}

	routes.UserRoute(r, authMiddleware)
	routes.ProductRoutes(r, authMiddleware)
	routes.CartRoutes(r, authMiddleware)
	routes.OrderRoutes(r, authMiddleware)

	r.Run()
}
