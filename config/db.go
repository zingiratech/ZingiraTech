package config

import (
	"context"
	"errors"

	"zingiratech/model"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Db(file string) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(file)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.New("error initializing app\n" + err.Error())
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.New("error initializing firestore client\n" + err.Error())
	}
	return client, nil
}

func AddUser(ctx context.Context, client *firestore.Client, user model.User) error {
	_, err := client.Collection("users").Doc(user.ID).Set(ctx, user)
	return err
}

func LogWasteCollection(ctx context.Context, client *firestore.Client, collection model.WasteCollection) error {
	_, _, err := client.Collection("waste_collections").Add(ctx, collection)
	return err
}

func GetUser(ctx context.Context, client *firestore.Client, userID string) (*model.User, error) {
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
