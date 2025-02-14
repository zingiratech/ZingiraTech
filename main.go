package main

import (
	"context"
	"fmt"
	"log"

	"zingiratech/config"
	"zingiratech/model"
)

func main() {
	ctx := context.Background()
	client, err := config.Db("zingiratech.json")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	user := model.User{
		ID:      "testuser123",
		Name:    "John Doe",
		Email:   "johndoe@example.com",
		Address: "Migosi",
	}
	err = config.AddUser(ctx, client, user)

	if err != nil {
		log.Fatalf("error adding user: %v", err)
	} else {
		fmt.Println("✅ User added successfully!")
	}

	retrievedUser, err := config.GetUser(ctx, client, user.ID)

	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	} else {
		fmt.Printf("🔍 Retrieved user: %+v\n", retrievedUser)
	}
}
