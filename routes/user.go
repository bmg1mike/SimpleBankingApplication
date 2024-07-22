package routes

import (
	"net/http"
	"simpleBankingApplication/models"

	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {
	var user models.User
	err:= context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid request",
		})
		return
	}

	err = user.SaveUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while saving the user"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"User saved successfully"})
}