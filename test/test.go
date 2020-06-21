package main

import (
	"encoding/json"
	"log"
)

func main() {
	JSONStruct := struct {
		FromLastDays int      `json:"from_last_days"`
		Usernames    []string `json:"usernames"`
		Type         string   `json:"type"`
		Limit        int      `json:"limit"`
		Offset       int      `json:"offset"`
	}{
		FromLastDays: 1,
		Usernames:    []string{"aicam", "aicam2"},
		Type:         "specific type if is needed",
		Limit:        10,
		Offset:       0,
	}
	js, _ := json.Marshal(JSONStruct)
	log.Print(string(js))
}
