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
			WrongRequestParameters(c, nil)
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
			WrongRequestParameters(c, nil)
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
		var bodyJson struct {
			Body     string `json:"body"`
			ImageUrl string `json:"image_url"`
		}
		err := c.BindJSON(&bodyJson)
		if err != nil {
			WrongRequestParameters(c, err)
			return
		}
		notifText := bodyJson.Body
		imageUrl := bodyJson.ImageUrl
		strings.ReplaceAll(notifText, "%USERNAME%", username)
		token, dbError := database.GetTokenByUsername(s.DB, username)
		if dbError != nil {
			FailedSqlCommand(c, dbError)
			return
		}
		message := FCMFuncs.GenerateMessage(imageUrl, notifText, title, token)
		messageID, fcmError := FCMFuncs.SendMessage(s.FCMApp, message)
		if fcmError != nil {
			FCMError(c, fcmError)
			database.AddFailedMessage(s.DB, username, fcmError.Error(), "single")
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

func (s *Server) addScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		score := c.Param("score")
		scoreInt, err := strconv.Atoi(score)
		if err != nil {
			WrongRequestParameters(c, err)
		}
		database.AddScoreModel(s.DB, username, scoreInt)
	}
}

func (s *Server) addMultipleScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var JSONStruct []struct {
			Username string `json:"username"`
			Score    int    `json:"score"`
		}
		err := c.BindJSON(&JSONStruct)
		if err != nil {
			WrongRequestParameters(c, err)
		}
		database.AddmultipleScoreModel(s.DB, JSONStruct)
		c.String(http.StatusOK, responses.ReturnSuccessedResponse("users updated successfully"))
	}
}

func (s *Server) sendMultipleNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var JSONStruct struct {
			Body     string   `json:"body"`
			Title    string   `json:"title"`
			ImageUrl string   `json:"image_url"`
			Users    []string `json:"users"`
		}
		err := c.BindJSON(&JSONStruct)
		if err != nil {
			WrongRequestParameters(c, err)
		}
		// check how many of notifications failed to send
		counter := 0
		for _, username := range JSONStruct.Users {
			token, err := database.GetTokenByUsername(s.DB, username)
			if err != nil {
				counter++
				continue
			}
			strings.ReplaceAll(JSONStruct.Body, "%USERNAME%", username)
			message := FCMFuncs.GenerateMessage(JSONStruct.ImageUrl, JSONStruct.Body, JSONStruct.Title, token)
			messageID, fcmError := FCMFuncs.SendMessage(s.FCMApp, message)
			if fcmError != nil {
				database.AddFailedMessage(s.DB, username, fcmError.Error(), "multiple")
				counter++
				continue
			}
			dbError := database.StoreMessageID(s.DB, messageID, username)
			if dbError != nil {
				FailedSqlCommand(c, dbError)
				return
			}
		}
		counterString := strconv.Itoa(counter)
		c.String(http.StatusOK, responses.ReturnSuccessedResponse(counterString+" number of notifications failed to send"))
	}
}

// limit of users retrieved is 20
func (s *Server) getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		offset, err := strconv.Atoi(c.Param("offset"))
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			WrongRequestParameters(c, err)
			return
		}
		users := database.GetUsersModel(s.DB, offset, limit)
		js, _ := json.Marshal(users)
		c.String(http.StatusOK, string(js))
	}
}

func (s *Server) removeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		err := database.RemoveUserModel(s.DB, username)
		if err != nil {
			FailedSqlCommand(c, err)
			return
		}
		c.String(http.StatusOK, responses.ReturnSuccessedResponse("user removed successfully"))
	}
}
