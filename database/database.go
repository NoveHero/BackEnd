package database

import (
	"gorm.io/driver/sqlite" // Your preferred database driver
	"gorm.io/gorm"

	"log"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil // Return nil on error
	}

	err = db.AutoMigrate(&User{}, &Story{}, &StoryMetaData{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
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
