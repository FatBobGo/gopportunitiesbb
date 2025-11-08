package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	// return nil
	return errors.New("not implemented")
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}