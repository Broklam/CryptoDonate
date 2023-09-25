package types

import (


	"gorm.io/gorm"
)

type PublicStreamers struct {
	gorm.Model
	Nickname string `json:"nickname" gorm:"unique"`
	Description string `json:"Description" gorm:"unique"`
	Twitch string `json:"Twitch" gorm:"unique"`
	Youtube string `json:"Youtube" gorm:"unique"`
	Telegram string `json:"Telegram" gorm:"unique"`
}

type PrivateStreamers struct {
	gorm.Model
	Nickname string 
	Name string 
	Email string
	Telegram string
	PassHash []byte 
	CreatedAt int
	LastLoginAt int
	BalanceBTC float32 
	BalanceTON float32
	BalanceETH float32
}

type Donations struct {
	gorm.Model
	FromNickname string
	ToStreamer string
	Message string
	Coin string
	AmountCoin float32
	Status int
	Accounted bool
}

type accounting struct {
	StreamerId int 
	InBtc float64
	InTon float64
	OutBtc float64
	OutTon float64
	OutEth float64
}
