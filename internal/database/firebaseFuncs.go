package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func AddFailedMessage(db *gorm.DB, username string, errorString string, typeOf string) {
	failedMessage := FailedMessages{
		Username:  username,
		ErrorSend: errorString,
		Type:      typeOf,
	}
	db.Save(&failedMessage)
}

func GetFailedMessages(db *gorm.DB, fromLastDays int, usernames []string, Type string, limit int, offset int) ([]FailedMessages, error) {
	var daysLimit time.Time
	var failedMessagesMatched []FailedMessages
	baseQuery := db

	if len(usernames) > 0 {
		for i, item := range usernames {
			if i == 0 {
				baseQuery = baseQuery.Where(&FailedMessages{Username: item})
				continue
			}
			baseQuery = baseQuery.Or(&FailedMessages{Username: item})
		}
	}
	if fromLastDays != 0 {
		daysLimit = time.Now().Add(-time.Hour * 24 * time.Duration(fromLastDays))
		baseQuery = baseQuery.Where(" created_at > ? ", daysLimit)
	}
	log.Print(daysLimit)
	if Type != "" {
		baseQuery = baseQuery.Where(&FailedMessages{Type: Type})
	}
	if limit == 0 {
		// default limit
		baseQuery = baseQuery.Limit(1000)
	} else {
		baseQuery = baseQuery.Limit(limit)
	}
	err := baseQuery.Offset(offset).Find(&failedMessagesMatched).Error
	return failedMessagesMatched, err
}
