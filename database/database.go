package database

import (
	// Your preferred database driver
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres" // Import the PostgreSQL driver
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	// Retrieve environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disabled TimeZone=UTC",
		host, user, password, dbname, port)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil // Return nil on error
	}

	// Run auto-migrations
	err = db.AutoMigrate(&User{}, &Story{}, &StoryMetaData{})
	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
		return nil
	}
	return db
}

// CloseDatabase closes the database connection if it's not nil.
func CloseDatabase(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting sql.DB: %v\n", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v\n", err)
		}
	}
}
