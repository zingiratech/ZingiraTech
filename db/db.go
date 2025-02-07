package db

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	home, err := os.Getwd()

	if err != nil {
		return err
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "JSON File Location")
	config := &firebase.Config{DatabaseURL: "Database URL"}
	app, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing databse: %v", err)
	}
	db.Client = client
	return nil
}

func FirebaseDB() *FireDB {
	return &fireDB
}
