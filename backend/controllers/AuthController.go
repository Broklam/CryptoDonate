package controllers

import (
	"log"
	"net/http"

	_ "github.com/Broklam/cryptodonate/backend/auth"
	"github.com/Broklam/cryptodonate/backend/db"
	"github.com/Broklam/cryptodonate/backend/encrypt"

	//"github.com/Broklam/cryptodonate/backend/encrypt"
	"github.com/Broklam/cryptodonate/backend/types"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`

	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user types.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := storage.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	err, isPasswordRight := encrypt.ComparePasswordAndHash(request.Password, user.PasswordHash)
	log.Println(isPasswordRight)
	log.Println(err)
	//tokenString, err := auth.GenerateJWT(user.Nickname, user.Username)
	/*if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}*/
	//context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
