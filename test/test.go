package main

import (
	"encoding/json"
	"log"
)

func main() {
	JSONStruct := struct {
		Body     string
		Title    string
		ImageUrl string
		Users    []string
	}{Body: "Your request body", Title: "Notif title", ImageUrl: "Image url", Users: []string{"aicam", "aicam2"}}
	js, _ := json.Marshal(JSONStruct)
	log.Print(string(js))
}
