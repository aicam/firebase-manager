package FCMFuncs

//import (
//	"github.com/aicam/notifServer/internal/database"
//	"github.com/jinzhu/gorm"
//)

//func getUsernameToken(db *gorm.DB, username string) string {
//	user := database.UsersFirebaseToken{Username: username}
//	err := db.Where(&database.UsersFirebaseToken{Username: username}).First(&user).RecordNotFound()
//	if err {
//		return ""
//	}
//	return user.Token
//}
