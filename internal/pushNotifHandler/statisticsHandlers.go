package pushNotifHandler

import "github.com/gin-gonic/gin"

func (s *Server) GetFailedMessagesByDate() gin.HandlerFunc {
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

	}
}
