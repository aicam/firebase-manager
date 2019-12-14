package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

func CheckUserNotExist(db *gorm.DB, username string) bool {
	return db.Where(&UsersData{Username: username}).Find(&UsersData{}).RecordNotFound()
}

func CheckUserTokenNotExist(db *gorm.DB, username string) bool {
	return db.Where(&UsersFirebaseToken{Username: username}).Find(&UsersFirebaseToken{}).RecordNotFound()
}

func CreateNewUserToken(db *gorm.DB, username string, token string) error {
	return db.Save(&UsersFirebaseToken{Username: username, Token: token}).Error
}

func UpdateUserToken(db *gorm.DB, username string, token string) error {
	userToken := UsersFirebaseToken{}
	sqlError := db.Where(&UsersFirebaseToken{Username: username}).First(&userToken).Error
	if sqlError != nil {
		return sqlError
	}
	if userToken.Token != token {
		db.Delete(&userToken)
		sqlError = db.Save(&UsersFirebaseToken{Token: token, Username: username}).Error
		if sqlError != nil {
			return sqlError
		}
	}
	return nil
}

func GetTokenByUsername(db *gorm.DB, username string) (string, error) {
	userToken := UsersFirebaseToken{}
	sqlError := db.Where(&UsersFirebaseToken{Username: username}).First(&userToken).Error
	if sqlError != nil {
		return "", sqlError
	}
	return userToken.Token, nil
}

func StoreMessageID(db *gorm.DB, messageId string, username string) error {
	err := db.Save(&MessageIDs{
		Username:  username,
		MessageId: messageId,
	}).Error
	return err
}

func AddScoreModel(db *gorm.DB, username string, score int) {
	user := UsersData{}
	db.Where(&UsersData{Username: username}).Find(&user)
	user.Score += score
	db.Save(&user)
}

func AddmultipleScoreModel(db *gorm.DB, bodyJSON []struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}) {
	for _, item := range bodyJSON {
		AddScoreModel(db, item.Username, item.Score)
	}
}

type UsersResponseData struct {
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Score     int        `json:"score"`
	Ban       bool       `json:"ban"`
	Token     string     `json:"token"`
}

func GetUsersModel(db *gorm.DB, offset int, limit int) []UsersResponseData {
	responseArray := []UsersResponseData{}
	users := []UsersData{}
	db.Order("score").Limit(limit).Offset(offset).Find(&users)
	userToken := UsersFirebaseToken{}
	for _, item := range users {
		db.Where(&UsersFirebaseToken{Username: item.Username}).First(&userToken)
		responseArray = append(responseArray, UsersResponseData{
			Username:  item.Username,
			CreatedAt: item.CreatedAt,
			DeletedAt: item.DeletedAt,
			Score:     item.Score,
			Ban:       item.Ban,
			Token:     userToken.Token,
		})
	}
	return responseArray
}
