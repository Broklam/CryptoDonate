package controllers

import (
	"net/http"

	storage "github.com/Broklam/cryptodonate/backend/db"

	"github.com/Broklam/cryptodonate/backend/types"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user types.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.PasswordHash); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := storage.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Status": "Succ"})
	//context.JSON(http.StatusCreated, gin.H{"user": user.ID, "email": user.Email, "username": user.Username})
}
