package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const (
	firebasConfigFile = "firebaseConfig.json"
	firebaseDBURL = "https://your-firebase-project.firebaseio.com"
)

var (
	ctx context.Context
	app *firebase.App
)


func main(){

ctx = context.Background()

opt := option.WithCredentialsFile(firebasConfigFile)

app, err := firebase.NewApp(ctx,, nil, opt)

if err != nil {
	log.Fatalf("Firebase initialization error: %v\n", err)
}


}
