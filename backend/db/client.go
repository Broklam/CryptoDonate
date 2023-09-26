package storage

import (
	//"github.com/Broklam/cryptodonate/backend/types"
	"log"

	"github.com/Broklam/cryptodonate/backend/types"
	"gorm.io/driver/sqlite" // Import SQLite driver
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error
var p string = "../db/data.db"

func Connect() {
	// Connect to SQLite database (data.db)
	Instance, dbError = gorm.Open(sqlite.Open(p), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	// Auto migrate the User model to the SQLite database
	Instance.AutoMigrate(&types.User{}, &types.Streamer{}, &types.Donation{}) //&types.Donation{})

	log.Println("Database Migration Completed!")
}

func updateStreamerBalances(db *gorm.DB, streamer *types.Streamer) {
	var balances struct {
		TotalBTC float64
		TotalETH float64
		TotalTON float64
	}
	db.Model(&types.Donation{}).Where("to_user = ?", streamer.Nickname).Select("SUM(amount) as total_btc, coin").Group("coin").Scan(&balances)

	streamer.BTCBalance = balances.TotalBTC
	streamer.ETHBalance = balances.TotalETH
	streamer.TONBalance = balances.TotalTON

	db.Save(streamer)
}
