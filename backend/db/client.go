package storage

import (
	//"github.com/Broklam/cryptodonate/backend/types"
	"log"
	"sync"
	"time"

	//"time"

	"github.com/Broklam/cryptodonate/backend/types"
	"gorm.io/driver/sqlite" // Import SQLite driver
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error
var p string = "../db/data.db"
var streamer types.Streamer
var coins = []string{"btc", "eth", "ton"}

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
	// Auto migrate types to the SQLite database
	Instance.AutoMigrate(&types.User{}, &types.Streamer{}, &types.Donation{}) //&types.Donation{})
	log.Println("Database Migration Completed!")
	err := updateSumBtc(Instance, "btc")
	if err != nil {
		log.Fatal(err)
	}

	err2 := updateSumEth(Instance, "ethc")
	if err2 != nil {
		log.Fatal(err)
	}

	err3 := updateSumTon(Instance, "ton")
	if err3 != nil {
		log.Fatal(err)
	}

}

func updateSumBtc(db *gorm.DB, coin string) error {
	// Raw SQL query to update SumCoin for a specific coin type
	query := `
        UPDATE streamers AS s
        SET btc_balance = (
            SELECT SUM(d.amount)
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        )
        WHERE EXISTS (
            SELECT 1
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        );
    `

	if err := db.Exec(query, coin, coin).Error; err != nil {
		return err
	}

	return nil
}

func updateSumEth(db *gorm.DB, coin string) error {
	// Raw SQL query to update SumCoin for a specific coin type
	query := `
        UPDATE streamers AS s
        SET eth_balance = (
            SELECT SUM(d.amount)
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        )
        WHERE EXISTS (
            SELECT 1
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        );
    `

	if err := db.Exec(query, coin, coin).Error; err != nil {
		return err
	}

	return nil
}
func updateSumTon(db *gorm.DB, coin string) error {
	// Raw SQL query to update SumCoin for a specific coin type
	query := `
        UPDATE streamers AS s
        SET ton_balance = (
            SELECT SUM(d.amount)
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        )
        WHERE EXISTS (
            SELECT 1
            FROM Donations AS d
            WHERE d.nickname = s.nickname
            AND d.coin = ?
        );
    `

	if err := db.Exec(query, coin, coin).Error; err != nil {
		return err
	}

	return nil
}

func BalanceSum() {
	ticker := time.NewTicker(10 * time.Second)

	// Create a WaitGroup to wait for all updates to finish.
	var wg sync.WaitGroup

	// Start a goroutine to run the updates every 10 seconds.
	go func() {
		for {
			select {
			case <-ticker.C:
				// Increment the WaitGroup counter for each update.
				wg.Add(3)

				// Run the updates in parallel Goroutines.
				go func() {
					defer wg.Done()
					err := updateSumBtc(Instance, "btc")
					if err != nil {
						log.Println(err)
					}
				}()

				go func() {
					defer wg.Done()
					err := updateSumEth(Instance, "ethc")
					if err != nil {
						log.Println(err)
					}
				}()

				go func() {
					defer wg.Done()
					err := updateSumTon(Instance, "ton")

					if err != nil {
						log.Println(err)
					}
					err2 := updateSumEth(Instance, "eth")
					if err2 != nil {
						log.Println(err)
					}
					err3 := updateSumBtc(Instance, "btc")
					if err3 != nil {
						log.Println(err)
					}
				}()
			}
		}
	}()

	// Wait for all updates to finish before exiting.
	wg.Wait()
}

func DeleteUserByID(username string) error {
	// Assuming you have a GORM database connection called "db" set up

	// First, check if the user exists
	var user types.User
	if err := Instance.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err // User not found
		}
		return err // Other database error
	}

	// If the user exists, delete them
	if err := Instance.Delete(&user).Error; err != nil {
		return err // Error occurred during deletion
	}

	return nil // User deleted successfully
}
