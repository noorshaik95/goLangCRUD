package main

import (
	"fmt"
	"goLandCRUD/config"
	"goLandCRUD/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	config.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":3000")
	if err != nil {
		panic(err)
	}
}
