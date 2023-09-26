package main

import (
	//"fmt"

	storage "github.com/Broklam/cryptodonate/backend/db"
	"github.com/Broklam/cryptodonate/backend/server"
	//"github.com/Broklam/cryptodonate/backend/db"
)

func main() {
	storage.Connect()
	storage.Migrate()
	router := server.InitRouter()
	router.Run(":8080")

}
