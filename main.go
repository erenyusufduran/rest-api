package main

import (
	"github.com/erenyusufduran/rest-api/db"
	"github.com/erenyusufduran/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
