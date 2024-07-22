package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", SayHello)
	server.POST("/make-payment", MakePayment)
	server.GET("/payment/:reference", GetPaymentByReference)
	server.POST("/add-user", AddUser)
}

func SayHello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
