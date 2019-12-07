package database

import (
	"github.com/jinzhu/gorm"
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

func GetUsersModel(db *gorm.DB, offset int) []UsersData {
	users := []UsersData{}
	db.Order("score").Limit(20).Offset(offset).Find(&users)
	return users
}
