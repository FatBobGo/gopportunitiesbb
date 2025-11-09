package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
	// return errors.New("not implemented")
}

func GetSQLite() (*gorm.DB, error) {
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
		return nil, err
	}
	return db, nil
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}