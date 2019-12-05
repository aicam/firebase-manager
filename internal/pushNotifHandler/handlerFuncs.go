package pushNotifHandler

import (
	"encoding/json"
	"github.com/aicam/notifServer/external/FCMFuncs"
	"github.com/aicam/notifServer/internal/database"
	"github.com/aicam/notifServer/internal/pushNotifHandler/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *Server) addUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		score := c.Param("score")
		scoreInt, err := strconv.Atoi(score)
		if err != nil || username == "" {
			WrongRequestParameters(c)
			return
		}
		user := database.UsersData{}
		notFound := s.DB.Where(&database.UsersData{Username: username}).First(&user).RecordNotFound()
		if !notFound {
			c.String(http.StatusOK, responses.ReturnSuccessedResponse("user is already exists"))
			return
		}
		s.DB.Save(&database.UsersData{Username: username, Score: scoreInt, Ban: false})
		res := responses.ResponseStructure{
			Status:    true,
			Data:      "User added",
			TimeStamp: time.Now().Unix(),
		}
		resJson, _ := json.Marshal(res)
		c.String(http.StatusOK, string(resJson))
	}
}

func (s *Server) setToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		token := c.Param("token")
		if username == "" || token == "" {
			WrongRequestParameters(c)
			return
		}
		if database.CheckUserNotExist(s.DB, username) {
			FailedLoadData(c)
			return
		}
		if database.CheckUserTokenNotExist(s.DB, username) {
			dbErr := database.CreateNewUserToken(s.DB, username, token)
			if dbErr != nil {
				FailedSqlCommand(c, dbErr)
				return
			}
		}
		dbErr := database.UpdateUserToken(s.DB, username, token)
		if dbErr != nil {
			FailedSqlCommand(c, dbErr)
		}
		c.String(http.StatusOK, responses.ReturnSuccessedResponse("user token updated"))
	}
}

func (s *Server) sendNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		title := c.Param("title")
		topic := c.Param("topic")
		var bodyJson struct {
			Body     string `json:"body"`
			ImageUrl string `json:"image_url"`
		}
		err := c.BindJSON(&bodyJson)
		if err != nil {
			WrongRequestParameters(c)
			return
		}
		notifText := bodyJson.Body
		imageUrl := bodyJson.ImageUrl
		strings.ReplaceAll(string(notifText), "%USERNAME%", username)
		token, dbError := database.GetTokenByUsername(s.DB, username)
		if dbError != nil {
			FailedSqlCommand(c, dbError)
			return
		}
		message := FCMFuncs.GenerateMessage(topic, imageUrl, notifText, title, token)
		messageID, fcmError := FCMFuncs.SendMessage(s.FCMApp, message)
		if fcmError != nil {
			FCMError(c, fcmError)
			return
		}
		dbError = database.StoreMessageID(s.DB, messageID, username)
		if dbError != nil {
			FailedSqlCommand(c, dbError)
			return
		}
		c.String(http.StatusOK, responses.ReturnSuccessedResponse(messageID))
	}
}

func (s *Server) addScore() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
