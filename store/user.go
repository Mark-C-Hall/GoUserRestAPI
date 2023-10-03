package store

import (
	"errors"
	"fmt"
	"sync"
	"user-api/util"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string
}

type inMemoryStore struct {
	users   []User
	userMap map[string]*User
	mutex   *sync.RWMutex
}

var store = inMemoryStore{
	users:   make([]User, 0),
	userMap: make(map[string]*User),
	mutex:   &sync.RWMutex{},
}

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

	u.ID = len(store.users) + 1
	store.users = append(store.users, *u)
	store.userMap[u.Username] = u

	return nil
}

func GetUserByUsername(username string) (User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user, exists := store.userMap[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return *user, nil
}

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

func DeleteUserByUsername(username string) error {
	return nil
}
