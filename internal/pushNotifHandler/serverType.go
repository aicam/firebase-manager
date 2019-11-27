package pushNotifHandler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server parts : database(mysql) - router
type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

// Here we create our new server
func NewServer() *Server {
	return &Server{
		DB:     nil,
		Router: mux.NewRouter(),
	}
}
