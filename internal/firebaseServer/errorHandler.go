package firebaseServer

import (
	"github.com/aicam/notifServer/internal/firebaseServer/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrongRequestParameters(c *gin.Context, err error) {
	c.String(http.StatusOK, responses.ReturnFailedResponse("request had wrong parameters , reasons : wrong number format , arguments you passed may be less or more, \n read the document and check your request: \n"+err.Error()))
}

func FailedLoadData(c *gin.Context) {
	c.String(http.StatusOK, responses.ReturnFailedResponse("It seems loading data failed , reasons: wrong username to take or set token , database is closed \n read the document and check your request"))
}

func FailedSqlCommand(c *gin.Context, err error) {
	c.String(http.StatusOK, responses.ReturnFailedResponse("sql operation failed , this may happen because your mysql server is not configured \n\n Error: \n"+err.Error()))
}

func FCMError(c *gin.Context, err error) {
	c.String(http.StatusOK, responses.ReturnFailedResponse("FCM (firebase Admin SDK) failed with error , check your google-service.json : "+err.Error()))
}
