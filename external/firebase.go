package external

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/jinzhu/gorm"
	"log"
	"os/exec"
	"strings"
)

// this package and file is used to communicate with firebase
// used only for linux
func SetGoogleServicePath(googleServicePath string) {
	cmd := exec.Command("export GOOGLE_APPLICATION_CREDENTIALS=\"" + googleServicePath + "\"")
	err := cmd.Run()
	if err != nil {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS couldn't set , this may happen because app has no permission")
	}
}

func initializeFirebase() *firebase.App {
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
		log.Fatalln(err)
		return "", err
	}
	respJson := map[string]string{}
	jsonError := json.Unmarshal([]byte(resp), &respJson)
	if jsonError != nil {
		log.Print("It seems parsing response json from firebase failed , it may happen if firebase api changes")
		return "", jsonError
	}
	return strings.Split(respJson["name"], "/")[3], nil
}

/* this method prepare message to send
parameters:
	title: title of notification to push
	body: body of notification to push
	imageUrl : url of image that will show in notification
	username: username that will notification send to
*/
func GenerateMessage(db *gorm.DB, topic string, imageUrl string, body string, title string, username string) *messaging.Message {
	return &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imageUrl,
		},
		Topic: topic,
		Token: getUsernameToken(db, username),
	}
}
