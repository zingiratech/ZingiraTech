package main

import (
	"fmt"
	"log"
	"net/http"

	"ZingiraTech/db"
)

const (
	firebasConfigFile = "firebaseConfig.json"
	firebaseDBURL     = "https://your-firebase-project.firebaseio.com"
)

func main() {
	err := db.FirebaseDB().Connect()

	if err != nil {
		log.Println(err)
	}

	router := mux.NewRouter()

	// TO DO : Defining API routes

	port := ":8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
