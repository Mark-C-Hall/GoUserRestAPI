package store

import (
	"errors"
	"fmt"
	"user-api/util"
)

// User represents a user with ID, username, email, and password fields.
// ID, username, and email fields are tagged to be included in JSON serialization, while the password field is excluded.
type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string
}

// CreateUser adds a new user to the in-memory store.
// It first hashes the password using a utility function, then assigns a unique ID to the user
// and finally adds the user to the userMap.
// Returns an error if the username already exists or if there's an error hashing the password.
func CreateUser(u *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, exists := store.userMap[u.Username]; exists {
		return errors.New("username already exists")
	}

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	u.Password = hashedPassword

	store.userCount++
	u.ID = store.userCount
	store.userMap[u.Username] = u

	return nil
}

// GetUserByUsername retrieves a user from the in-memory store by username.
// Returns the user and an error if the user is not found.
func GetUserByUsername(username string) (User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user, exists := store.userMap[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return *user, nil
}

// UpdateUser updates the details of an existing user in the in-memory store.
// It updates only the provided fields: email and password. For updating the password,
// it first hashes the new password and then replaces the old one.
// Returns an error if the user is not found or if there's an error hashing the password.
func UpdateUser(u *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	storeUser, exists := store.userMap[u.Username]
	if !exists {
		return errors.New("user not found")
	}

	if u.Email != "" {
		storeUser.Email = u.Email
	}

	if u.Password != "" {
		hashedPassword, err := util.HashPassword(u.Password)
		if err != nil {
			return errors.New("failed to hash password")
		}
		storeUser.Password = hashedPassword
	}

	return nil
}

// DeleteUserByUsername removes a user from the in-memory store by username.
// Returns an error if the user is not found.
func DeleteUserByUsername(username string) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	_, exists := store.userMap[username]
	if !exists {
		return errors.New("user not found")
	}
	delete(store.userMap, username)

	return nil
}
