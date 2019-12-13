package database

import (
	"github.com/jinzhu/gorm"
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

func GetFailedMessages(db *gorm.DB, fromLastDays int, usernames []string, Type string, limit int, offset int) []FailedMessages {
	var daysLimit time.Time
	var query string
	var args []string
	var failedMessagesMatched []FailedMessages
	query = " type = ? "
	args = append(args, Type)
	// we add new limitations by their value
	// usernames will add by query conditions
	if len(usernames) > 0 {
		query += "AND ("
		for index, item := range usernames {
			query += " username = ? "
			if index != len(usernames)-1 {
				query += "OR"
			}
			args = append(args, item)
		}
		query += ")"
	}
	if fromLastDays == 0 {
		daysLimit = time.Now().Add(-time.Hour * 24 * time.Duration(fromLastDays))
		query += " AND created_at > ? "
		args = append(args, daysLimit.String())
	}
	if limit == 0 {
		limit = 100
	}
	db.Where(query, args).Limit(limit).Offset(offset).Find(&failedMessagesMatched)
	return failedMessagesMatched
}

func getFailedMessagesType1(db *gorm.DB, typeOf string) []FailedMessages {
	var failedMessages []FailedMessages
	db.Where(" type = ? ", typeOf).Limit(100).Find(&failedMessages)
	return failedMessages
}

func getFailedMessagesType2(db *gorm.DB, typeOf string, offset int) []FailedMessages {
	var failedMessages []FailedMessages
	db.Where(" type = ? ", typeOf).Limit(100).Offset(offset).Find(&failedMessages)
	return failedMessages
}

func getFailedMessagesType3(db *gorm.DB, typeOf string, limit int) []FailedMessages {
	var failedMessages []FailedMessages
	db.Where(" type = ? ", typeOf).Limit(limit).Find(&failedMessages)
	return failedMessages
}

func getFailedMessagesType4(db *gorm.DB, typeOf string, limit int, offset int) []FailedMessages {
	var failedMessages []FailedMessages
	db.Where(" type = ? ", typeOf).Limit(limit).Offset(offset).Find(&failedMessages)
	return failedMessages
}
