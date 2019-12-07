package pushNotifHandler

import (
	firebase "firebase.google.com/go"
	"github.com/aicam/notifServer/external/FCMFuncs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Server parts : database(mysql) - router
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
	FCMApp *firebase.App
}

// Here we create our new server
func NewServer() *Server {
	router := gin.Default()
	// here we opened cors for all
	router.Use(cors.Default())
	return &Server{
		DB:     nil,
		Router: router,
		FCMApp: FCMFuncs.InitializeFirebase(),
	}
}
