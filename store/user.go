package store

import (
	"errors"
	"sync"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string
}

type inMemoryStore struct {
	users   []User
	userMap map[string]User
	mutex   *sync.RWMutex
}

var store = inMemoryStore{
	users:   make([]User, 0),
	userMap: make(map[string]User),
	mutex:   &sync.RWMutex{},
}

func CreateUser(u *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, exists := store.userMap[u.Username]; exists {
		return errors.New("username already exists")
	}

	u.ID = len(store.users) + 1
	store.users = append(store.users, *u)
	store.userMap[u.Username] = *u

	return nil
}

func GetUserByUsername(username string) (User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user, exists := store.userMap[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func UpdateUser(u *User) error {
	return nil
}

func DeleteUserByUsername(username string) error {
	return nil
}
