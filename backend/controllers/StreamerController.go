package controllers

import (
	"net/http"

	storage "github.com/Broklam/cryptodonate/backend/db"

	"github.com/Broklam/cryptodonate/backend/types"
	"github.com/gin-gonic/gin"
)

type StreamerRequest struct {
	Nickname string `json:"Nickname"`
}

// Search streamer by name controller
func FindStreamerByName(context *gin.Context) {
	var request StreamerRequest
	var streamer types.Streamer
	record := storage.Instance.First(&streamer, request.Nickname)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"streamer": streamer})
}

func RegisterFullStreamer(context *gin.Context) {
	var streamer types.Streamer

	// You should bind the request body to the streamer struct.
	if err := context.ShouldBindJSON(&streamer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the streamer record in the database.
	if err := storage.Instance.Create(&streamer).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message or the created streamer data.
	context.JSON(http.StatusCreated, "ok")
}
