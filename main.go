package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"zingiratech/config"
)

func main() {
	app, err := config.Db()
	if err != nil {
		log.Fatal("error getting firebase app: ", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal("error initializing firestore client")
	}

	defer client.Close()

	port := ":8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
