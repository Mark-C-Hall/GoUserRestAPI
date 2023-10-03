package store

import "sync"

// inMemoryStore is an in-memory data structure used to store and manage user data.
type inMemoryStore struct {
	userMap           map[string]*User // A map to store user data by username as the key
	userCount         int              // Count of total users, used to assign unique IDs
	mutex             *sync.RWMutex    // Mutex to ensure concurrent safe access to the userMap
	blacklistedTokens map[string]bool  // A map to store blacklisted tokens
}

// store is the in-memory database instance.
var store = inMemoryStore{
	userMap:           make(map[string]*User),
	mutex:             &sync.RWMutex{},
	blacklistedTokens: make(map[string]bool),
}
