package firebaseServer

import (
	"encoding/json"
	"github.com/aicam/notifServer/internal/database"
	"github.com/aicam/notifServer/internal/firebaseServer/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getFailedMessagesByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataFormat struct {
			FromLastDays int      `json:"from_last_days"`
			Usernames    []string `json:"usernames"`
			Type         string   `json:"type"`
			Limit        int      `json:"limit"`
			Offset       int      `json:"offset"`
		}
		err := c.BindJSON(&dataFormat)
		if err != nil {
			FCMError(c, err)
		}
		response, dbError := database.GetFailedMessages(s.DB, dataFormat.FromLastDays, dataFormat.Usernames, dataFormat.Type, dataFormat.Limit, dataFormat.Offset)
		if dbError != nil {
			FailedSqlCommand(c, dbError)
		}
		responseJSON, err := json.Marshal(response)
		c.String(http.StatusOK, responses.ReturnSuccessedResponse(string(responseJSON)))
	}
}
