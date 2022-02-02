package onlyugobind

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbError error

func OpenDatabase(payload string) (string, error) {
	db, dbError = gorm.Open(sqlite.Open(payload+"onlyu.db"), &gorm.Config{})
	if dbError != nil {
		return "Error", dbError
	}

	// Auto megrate
	db.AutoMigrate(&User{})

	return "Opened", nil
}
