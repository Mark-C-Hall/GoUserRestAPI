package store

// AddTokenToBlacklist adds a given JWT token to the blacklist.
// Once a token is blacklisted, it's considered invalid for further authentications.
//
// Parameters:
// - token: The JWT token string to be blacklisted.
func AddTokenToBlacklist(token string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.blacklistedTokens[token] = true
}

// IsTokenBlacklisted checks if a given JWT token is in the blacklist.
//
// Parameters:
// - token: The JWT token string to be checked.
//
// Returns:
// - true if the token is found in the blacklist; false otherwise.
func IsTokenBlacklisted(token string) bool {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	_, exists := store.blacklistedTokens[token]
	return exists
}
