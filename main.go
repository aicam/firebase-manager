package main

import (
	"github.com/aicam/notifServer/internal/database"
	"github.com/aicam/notifServer/internal/pushNotifHandler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

const DatabaseConnectionString  = "aicam:021021ali@tcp(127.0.0.1:3306)/shop_data?charset=utf8mb4&parseTime=True"

func main(){
	s := pushNotifHandler.NewServer()
	db := database.MakeMigrations(DatabaseConnectionString)
	s.DB = db
	err := http.ListenAndServe("0.0.0.0:4300", s.Router)
	if err != nil {
		log.Print(err)
	}
}