package types

import (
	"github.com/Broklam/cryptodonate/backend/encrypt"
)

type User struct {
	Username     string `gorm:"unique" json:"Username"`
	Nickname     string `gorm:"primaryKey; unique" json:"Nickname"`
	Password     string
	PasswordHash string     `json:"PasswordHash"`
	Role         uint8      `json:"Role"`
	Streamers    []Streamer `gorm:"foreignKey:Nickname"`
}

// Define the Streamer struct
type Streamer struct {
	Nickname    string     `gorm:"unique; primaryKey" json:"Nickname"`
	Description string     ` json:"Description"`
	BTCBalance  float64    `json:"BTCBalance"`
	ETHBalance  float64    `json:"ETHBalance"`
	TONBalance  float64    `json:"TONBalance"`
	Donations   []Donation `gorm:"foreignKey:Nickname" json:"Donations`
	Telegram    string     ` json:"Telegram"`
	Twitch      string     `json:"Twitch`
	Youtube     string     ` json:"Youtube`
}

// Define the Donation struct
type Donation struct {
	From          string  `gorm:"primaryKey" json:"From"`
	Nickname      string  `json:"ToStreamer"`
	Message       string  `json:"Message"`
	Coin          string  `json:"Coin"`
	Amount        float64 `json:"Amount"`
	PublicWallet  string  `gorm:"unique" json:"PublicWallet"`
	PrivateWallet string  `gorm:"unique" json:"PrivateWallet"`
	Status        uint8   `json:"Status"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := encrypt.GenerateFromPassword(password)
	if err != nil {
		return err
	}
	user.PasswordHash = string(bytes)
	return nil
}
