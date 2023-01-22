package main

import (
	"fmt"

	"go-shop-api/configs"
	"go-shop-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	configs.ConnectDB()

	user, err := configs.DeleteUser(configs.DB, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	routes.UserRoute(r)

	r.Run()
}
