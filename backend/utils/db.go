package utils

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
	"log"
	"sync"

	//"golang.org/x/crypto/bcrypt" 
)

type DB struct {
	*sql.DB
}

// createDB initializes a new database connection
func createDB() *DB {
	dbAddress := os.Getenv("DB_ADDRESS")
	if dbAddress == "" {
		panic("DB_ADDRESS is not set properly")
	}
	log.Println("Connecting to database at", dbAddress)
	db, err := sql.Open("postgres", dbAddress)
	if err != nil {
		log.Println("Error connecting to database:", err.Error())
		return nil
	}
	log.Println("Connected to database")
	return &DB{db}
}

var dbOnce sync.Once
var appDB *DB

// GetDB returns a singleton instance of the database connection
func GetDB() *DB {
	dbOnce.Do(func() {
		appDB = createDB()
	})
	return appDB
}
