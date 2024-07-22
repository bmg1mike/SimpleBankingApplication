package main

import (
	"simpleBankingApplication/db"
	"simpleBankingApplication/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	
	db.InitDb()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}