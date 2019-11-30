package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// Synchronize database
func MakeMigrations(connectionString string) *gorm.DB {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Print("mysql connection error : ", err)
	}
	db.AutoMigrate(&UsersData{})
	db.AutoMigrate(&UsersFirebaseToken{})
	db.AutoMigrate(&MessageIDs{})
	return db
}
