package main

import (
	"github.com/KritapasSuwannawin/Go-Project/db"
	"github.com/KritapasSuwannawin/Go-Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
