package main

import (
	"github.com/MuShaf-NMS/Skyshi-Test/config"
	"github.com/MuShaf-NMS/Skyshi-Test/database"
	"github.com/MuShaf-NMS/Skyshi-Test/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config := config.GetConfig()
	db := database.CreateConnection(config)
	gin.SetMode(gin.ReleaseMode)
	defer database.CloseConnection(db)
	server := gin.Default()
	server.Use(cors.Default())
	router.InitializeRoute(server, db, *config)
	server.Run(":3030")
}
