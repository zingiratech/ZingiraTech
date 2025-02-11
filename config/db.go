package config

import (
	"context"
	"log"

	"zingiratech/model"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Db() (*firebase.App, error) {
	opt := option.WithCredentialsFile("zingiratech.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return nil, err
	}
	return app, nil
}

func AddUser(ctx context.Context, client *firestore.Client, user model.User) error {
	_, err := client.Collection("users").Doc(user.ID).Set(ctx, user)
	return err
}

func logWasteCollection(ctx context.Context, client *firestore.Client, collection model.WasteCollection) error {
	_, _, err := client.Collection("waste_collections").Add(ctx, collection)
	return err
}

func letUser(ctx context.Context, client *firestore.Client, userID string) (*model.User, error) {
	doc, err := client.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
