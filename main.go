package main

import (
	"github.com/aicam/notifServer/external/FCMFuncs"
	_ "github.com/aicam/notifServer/external/FCMFuncs"
	"github.com/aicam/notifServer/internal/database"
	"github.com/aicam/notifServer/internal/pushNotifHandler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

const DatabaseConnectionString = "aicam:021021ali@tcp(mysql-db:3306)/firebase?charset=utf8mb4&parseTime=True"

func main() {
	args := os.Args
	var googleServicePath string //google-service.json path
	if len(args) <= 1 {
		googleServicePath = "/home/ali/go/src/github.com/aicam/notifServer/libs/google-services.json"
		log.Print("google-service.json set automatically")
	} else {
		googleServicePath = args[1]
		log.Print("google-service.json set " + googleServicePath)
	}
	FCMFuncs.SetGoogleServicePath(googleServicePath)
	// initialize new server with db and router
	s := pushNotifHandler.NewServer()
	// initialize database
	db := database.MakeMigrations(DatabaseConnectionString)
	s.DB = db
	s.Routes()
	err := http.ListenAndServe("0.0.0.0:4300", s.Router)
	if err != nil {
		log.Print(err)
	}
}
