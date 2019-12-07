package FCMFuncs

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"log"
	"os"
	"strings"
)

// this package and file is used to communicate with firebase
// used only for linux
func SetGoogleServicePath(googleServicePath string) {
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", googleServicePath)
	if err != nil {
		log.Print("GOOGLE_APPLICATION_CREDENTIALS couldn't set , this may happen because app has no permission \r\n", err)
	}
}

func InitializeFirebase() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("error initializing app: %v\n", err)
	}
	return app
}

func SendMessage(app *firebase.App, message *messaging.Message) (string, error) {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Print("error getting Messaging client: %v\n", err)
	}
	// Send a message to the device corresponding to the provided
	// registration token.
	resp, err := client.Send(ctx, message)
	if err != nil {
		log.Print(err)
		return "", err
	}
	// document of firebase said response is json with field name but it isn't :)
	//respJson := map[string]string{}
	//log.Print(resp)
	//jsonError := json.Unmarshal([]byte(resp), &respJson)
	//if jsonError != nil {
	//	log.Print("It seems parsing response json from firebase failed , it may happen if firebase api changes")
	//	return "", jsonError
	//}
	return strings.Split(resp, "/")[3], nil
}

/* this method prepare message to send
parameters:
	title: title of notification to push
	body: body of notification to push
	imageUrl : url of image that will show in notification
	username: username that will notification send to
*/
func GenerateMessage(imageUrl string, body string, title string, token string) *messaging.Message {
	return &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imageUrl,
		},
		Token: token,
	}
}
