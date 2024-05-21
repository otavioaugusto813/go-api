package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
	// return errors.New("fake Error")
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}
