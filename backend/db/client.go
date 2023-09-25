package storage

import (
	//"github.com/Broklam/cryptodonate/backend/types"
	"log"

	"github.com/Broklam/cryptodonate/backend/types"
	"gorm.io/driver/sqlite" // Import SQLite driver
	"gorm.io/gorm"
)

var PublicStreamers types.PublicStreamers
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
	Instance.AutoMigrate(&PublicStreamers)
	log.Println("Database Migration Completed!")
}
