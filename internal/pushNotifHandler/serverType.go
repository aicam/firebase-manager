package pushNotifHandler

import (
	firebase "firebase.google.com/go"
	"github.com/aicam/notifServer/external/FCMFuncs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server parts : database(mysql) - router
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	FCMApp *firebase.App
}

// Here we create our new server
func NewServer() *Server {
	return &Server{
		DB:     nil,
		Router: mux.NewRouter(),
		FCMApp: FCMFuncs.InitializeFirebase(),
	}
}
