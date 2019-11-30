package database

import "github.com/jinzhu/gorm"

type UsersFirebaseToken struct {
	gorm.Model
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UsersData struct {
	gorm.Model
	Username string `json:"username"`
	Score    int    `json:"score"`
	Ban      bool   `json:"ban"`
}

type MessageIDs struct {
	gorm.Model
	Username  string `json:"username"`
	MessageId string `json:"message_id"`
}
