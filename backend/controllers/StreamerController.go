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

type StreamerRegistrationRequest struct {
	Nickname    string
	Name        string
	Email       string
	Youtube     string
	Telegram    string
	Password    string
	PassHash    string
	CreatedAt   int
	LastLoginAt int
	BalanceBTC  float32
	BalanceTON  float32
	BalanceETH  float32
}

// Search streamer by name controller
func FindStreamerByName(context *gin.Context) {
	var request StreamerRequest
	var streamer types.PublicStreamers
	record := storage.Instance.First(&streamer, request.Nickname)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"streamer": streamer})
}
