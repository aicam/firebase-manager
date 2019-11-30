package responses

import (
	"encoding/json"
	"time"
)

type ResponseStructure struct {
	Status    bool   `json:"status"`
	Data      string `json:"data"`
	TimeStamp int64  `json:"time_stamp"`
}

func ReturnFailedResponse(data string) []byte {
	res := ResponseStructure{
		Status:    false,
		Data:      data,
		TimeStamp: time.Now().Unix(),
	}
	resJson, _ := json.Marshal(res)
	return resJson
}

func ReturnSuccessedResponse(data string) []byte {
	res := ResponseStructure{
		Status:    true,
		Data:      data,
		TimeStamp: time.Now().Unix(),
	}
	resJson, _ := json.Marshal(res)
	return resJson
}
