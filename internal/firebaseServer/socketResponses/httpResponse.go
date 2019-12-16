package socketResponses

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Status         bool   `json:"status"`
	ErrorMessage   string `json:"error_message"`
	ServerResponse string `json:"server_response"`
}

func ErrorOnConnection(w http.ResponseWriter, err error) {
	response := httpResponse{
		Status:         false,
		ErrorMessage:   err.Error(),
		ServerResponse: "failed to initial socket connection",
	}
	responseJSON, _ := json.Marshal(response)
	_, _ = w.Write(responseJSON)
}

func ErrorOnReadMessage(w http.ResponseWriter, err error) {
	response := httpResponse{
		Status:         false,
		ErrorMessage:   err.Error(),
		ServerResponse: "Reading from socket failed",
	}
	responseJSON, _ := json.Marshal(response)
	_, _ = w.Write(responseJSON)
}

func ErrorOnSendMessage(w http.ResponseWriter, err error) {
	response := httpResponse{
		Status:         false,
		ErrorMessage:   err.Error(),
		ServerResponse: "Writing on socket failed",
	}
	responseJSON, _ := json.Marshal(response)
	_, _ = w.Write(responseJSON)
}
