package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UsersFirebaseToken struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type BackupTokens struct {
	gorm.Model
	Username string `json:"username"`
	Token string `json:"token"`
}

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
