package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func Tick(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "u did it finally"})
}