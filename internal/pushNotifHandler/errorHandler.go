package pushNotifHandler

import (
	"github.com/aicam/notifServer/internal/pushNotifHandler/responses"
	"net/http"
)

func WrongRequestParameters(writer http.ResponseWriter) {
	_, _ = writer.Write(responses.ReturnFailedResponse("request had wrong parameters , reasons : wrong number format , arguments you passed may be less or more, \n read the document and check your request"))
}

func FailedLoadData(writer http.ResponseWriter) {
	_, _ = writer.Write(responses.ReturnFailedResponse("It seems loading data failed , reasons: wrong username to take or set token , database is closed \n read the document and check your request"))
}

func FailedSqlCommand(writer http.ResponseWriter, err error) {
	_, _ = writer.Write(responses.ReturnFailedResponse("sql operation failed , this may happen because your mysql server is not configured \n\n Error: \n" + err.Error()))
}

func FCMError(writer http.ResponseWriter, err error) {
	_, _ = writer.Write(responses.ReturnFailedResponse("FCM (firebase Admin SDK) failed with error , check your google-service.json : " + err.Error()))
}
