package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// Synchronize database
func MakeMigrations(connectionString string) *gorm.DB {
	db, err := gorm.Open(connectionString)
	if err != nil {
		log.Print(err)
	}
	db.AutoMigrate(&UsersFirebaseToken{})
	db.AutoMigrate(&BackupTokens{})
	return db
}
