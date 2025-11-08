package config

import (
	"os"

	"github.com/nathan/gopportunitiesbb/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	logger.Info("Initializing SQLite database connection...")

	// Check if the database file exists, if not create it
	dbPath := "./db/database.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("datbase file %s not found, creating...", dbPath)
		// create the database file and directory if not exist
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("failed to create database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
		logger.Info("database file created successfully")
	} else {
		logger.Info("database file already exists")
	}

	db, err := gorm.Open(sqlite.Open("./db/database.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto-migration error: %v", err)
		return nil, err
	}

	return db, nil
}