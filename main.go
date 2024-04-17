package main

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //returns a pointer to the Engine (server)

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080 for developement
}
