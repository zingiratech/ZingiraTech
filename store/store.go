package store

import (
	"context"
	"errors"
	"time"

	"ZingiraTech/db"
	"ZingiraTech/model"
)

type Store struct {
	*db.FireDB
}

// NewStore returns a Store.
func NewStore() *Store {
	d := db.FirebaseDB()
	return &Store{
		FireDB: d,
	}
}

// Create a new User object
func (s *Store) Create(b *model.User) error {
	if err := s.NewRef("users/"+b.ID.String()).Set(context.Background(), b); err != nil {
		return err
	}
	return nil
}

// Get user by ID
func (s *Store) GetByID(userID string) (*model.User, error) {
	user := &model.User{}
	if err := s.NewRef("users/"+userID).Get(context.Background(), user); err != nil {
		return nil, err
	}
	if user.Username == "" {
		return nil, nil
	}
	return user, nil
}

// Soft delete a user (marks as deleted)
func (s *Store) Delete(userID string, requestingUserID string) error {
	if requestingUserID != userID && !s.isAdmin(requestingUserID) {
		return errors.New("unauthorized deletion attempt")
	}

	ref := s.NewRef("users/" + userID)

	// transaction to checks if the user exists before deleting
	err := ref.Transaction(context.Background(), func(node db.DataNode) (interface{}, error) {
		user := &model.User{}
		if err := node.Unmarshal(user); err != nil {
			return nil, err
		}
		if user.Username == "" {
			return nil, errors.New("user does not exist")
		}

		// Mark as deleted instead of removing
		now := time.Now()
		user.DeletedAt = &now
		return user, nil
	})

	return err
}

// Update a user with given fields
func (s *Store) Update(userID string, updates map[string]interface{}) error {
	return s.NewRef("users/"+userID).Update(context.Background(), updates)
}

func (s *Store) isAdmin(userID string) bool {
	// TODO: Implement actual admin verification
	return userID == "admin-user-id"
}
