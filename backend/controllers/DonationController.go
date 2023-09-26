package controllers

import (
	"net/http"

	storage "github.com/Broklam/cryptodonate/backend/db"
	"github.com/Broklam/cryptodonate/backend/types"
	"github.com/gin-gonic/gin"
)

func CreateDonation(context *gin.Context) {
	var Donation types.Donation
	if err := context.ShouldBindJSON(&Donation); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := storage.Instance.Create(&Donation)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	//Generate wallet
	context.JSON(http.StatusCreated, gin.H{"From": Donation.From, "To": Donation.Nickname, "Coin": Donation.Coin, "Amount": Donation.Amount, "Address": Donation.PublicWallet, "Status": "Created"})
	//context.JSON(http.StatusCreated, gin.H{"Status": "Donation was created", "From": Donation.From, "To": Donation.Nickname, "Coin": Donation.Coin, "Amount": Donation.Amount, "Address": Donation.PublicWallet})
}
