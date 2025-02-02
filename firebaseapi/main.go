package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const (
	firebasConfigFile = "firebaseConfig.json"
	firebaseDBURL     = "https://your-firebase-project.firebaseio.com"
)

var (
	ctx context.Context
	app *firebase.App
)

func main() {

	ctx = context.Background()

	opt := option.WithCredentialsFile(firebasConfigFile)

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatalf("Firebase initialization error: %v\n", err)
	}

	client, err := app.DatabaseWithURL(ctx, firebaseDBURL)

	if err != nil {
		log.Fatalf("Firestore initialization error: %v\n", err)
	}

	router := mux.NewRouter()

	// TO DO : Defining API routes



	port := ":8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
