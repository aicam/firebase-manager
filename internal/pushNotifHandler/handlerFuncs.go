package pushNotifHandler

import (
	"encoding/json"
	"github.com/aicam/notifServer/external/FCMFuncs"
	"github.com/aicam/notifServer/internal/database"
	"github.com/aicam/notifServer/internal/pushNotifHandler/responses"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *Server) addUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		username := vars["username"]
		score := vars["score"]
		scoreInt, err := strconv.Atoi(score)
		if err != nil || username == "" {
			WrongRequestParameters(writer)
			return
		}
		s.DB.Save(&database.UsersData{Username: username, Score: scoreInt, Ban: false})
		res := responses.ResponseStructure{
			Status:    true,
			Data:      "User added",
			TimeStamp: time.Now().Unix(),
		}
		resJson, _ := json.Marshal(res)
		_, _ = writer.Write(resJson)
	}
}

func (s *Server) setToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		username := vars["username"]
		token := vars["token"]
		if username == "" || token == "" {
			WrongRequestParameters(writer)
			return
		}
		if database.CheckUserNotExist(s.DB, username) {
			FailedLoadData(writer)
			return
		}
		if database.CheckUserTokenNotExist(s.DB, username) {
			dbErr := database.CreateNewUserToken(s.DB, username, token)
			if dbErr != nil {
				FailedSqlCommand(writer, dbErr)
				return
			}
		}
		dbErr := database.UpdateUserToken(s.DB, username, token)
		if dbErr != nil {
			FailedSqlCommand(writer, dbErr)
		}
		_, _ = writer.Write(responses.ReturnSuccessedResponse("user token updated"))
	}
}

func (s *Server) sendNotification() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		username := vars["username"]
		title := vars["title"]
		topic := vars["topic"]
		body, _ := ioutil.ReadAll(request.Body)
		var bodyJson map[string]string
		err := json.Unmarshal(body, &bodyJson)
		if err != nil {
			WrongRequestParameters(writer)
			return
		}
		notifText := bodyJson["body"]
		imageUrl := bodyJson["image_url"]
		strings.ReplaceAll(string(notifText), "%USERNAME%", username)
		token, dbError := database.GetTokenByUsername(s.DB, username)
		if dbError != nil {
			FailedSqlCommand(writer, dbError)
			return
		}
		message := FCMFuncs.GenerateMessage(topic, imageUrl, notifText, title, token)
		messageID, fcmError := FCMFuncs.SendMessage(s.FCMApp, message)
		if fcmError != nil {
			FCMError(writer, fcmError)
			return
		}
		dbError = database.StoreMessageID(s.DB, messageID, username)
		if dbError != nil {
			FailedSqlCommand(writer, dbError)
			return
		}
		_, _ = writer.Write(responses.ReturnSuccessedResponse(messageID))
	}
}
