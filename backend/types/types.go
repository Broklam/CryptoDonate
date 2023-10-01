package types

import (
	"time"

	"github.com/Broklam/cryptodonate/backend/encrypt"
)

type User struct {
	Username     string     `gorm:"unique" json:"Username"`
	Nickname     string     `gorm:"primaryKey; unique" json:"Nickname"`
	Password     string     `json:"Nickname"`
	PasswordHash string     `json:"PasswordHash"`
	Role         uint8      `json:"Role"`
	Streamers    []Streamer `gorm:"foreignKey:Nickname"`
	CreatedAt    time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time  `json:"updaed_at" gorm:"column:updated_at"`
}

// Define the Streamer struct
type Streamer struct {
	Nickname    string     `gorm:"unique; primaryKey" json:"Nickname"`
	Description string     ` json:"Description"`
	BTCBalance  float64    `json:"BTCBalance"`
	ETHBalance  float64    `json:"ETHBalance"`
	TONBalance  float64    `json:"TONBalance"`
	Donations   []Donation `gorm:"foreignKey:Nickname" json:"Donations"`
	Telegram    string     ` json:"Telegram"`
	Twitch      string     `json:"Twitch"`
	Youtube     string     ` json:"Youtube"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

// Define the Donation struct
type Donation struct {
	From          string    `json:"From"`
	Nickname      string    `json:"ToStreamer"`
	Message       string    `json:"Message"`
	Coin          string    `json:"Coin"`
	Amount        float64   `json:"Amount"`
	PublicWallet  string    `json:"PublicWallet"`
	PrivateWallet string    `json:"PrivateWallet"`
	Address       string    `json:"Address"`
	Status        uint8     `json:"Status"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updaed_at" gorm:"column:updated_at"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := encrypt.GenerateFromPassword(password)
	if err != nil {
		return err
	}
	user.PasswordHash = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) (err error, result bool) {
	result, err = encrypt.ComparePasswordAndHash(providedPassword, user.PasswordHash)
	if err != nil {
		return err, false
	}
	return err, result
}
