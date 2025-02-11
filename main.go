package main

import (
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

	defer app.Close()

	port := ":8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
