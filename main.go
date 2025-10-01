package main

import (
	"fmt"
	"goLandCRUD/config"
	"goLandCRUD/logger"
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
		logger.Error("Error starting the server: ", err)
		panic(err)
	}
	logger.Info("Server Started Successfully!!")
}
