package external

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"log"
	"os/exec"
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

func SendMessage(app *firebase.App, registerationToken string, message *messaging.Message) {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Print("error getting Messaging client: %v\n", err)
	}
	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
}
