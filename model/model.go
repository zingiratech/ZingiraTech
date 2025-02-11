package model

import "time"

type User struct {
	ID      string `firestore:"id"`
	Name    string `firestore:"name"`
	Email   string `firestore:"email"`
	Address string `firestore:"address"`
}

type WasteCollection struct {
	UserID         string    `firestore:"user_id"`
	CollectionDate time.Time `firestore:"collection_date"`
	WasteType      string    `firestore:"waste_type"`
	Quantity       float64   `firestore:"quantity"`
}
