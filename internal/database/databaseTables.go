package database

import "github.com/jinzhu/gorm"

type UsersFirebaseToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type BackupTokens struct {
	gorm.Model
	Username string `json:"username"`
	Token    string `json:"token"`
}

type MessageIDs struct {
	gorm.Model
	Username  string `json:"username"`
	MessageId string `json:"message_id"`
}
