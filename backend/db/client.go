package storage

import (
   //"github.com/Broklam/cryptoDonate/backend/types"
    "log"
    "gorm.io/driver/sqlite" // Import SQLite driver
    "gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error
var p string = "./storage/data.db"
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
    //Instance.AutoMigrate(&types.User{})
    log.Println("Database Migration Completed!")
}
