package database

import "github.com/jinzhu/gorm"

func AddFailedMessage(db *gorm.DB, username string, errorString string, typeOf string) {
	failedMessage := FailedMessages{
		Username:  username,
		ErrorSend: errorString,
		Type:      typeOf,
	}
	db.Save(&failedMessage)
}
