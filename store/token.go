package store

func AddTokenToBlacklist(token string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.blacklistedTokens[token] = true
}

func IsTokenBlacklisted(token string) bool {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	_, exists := store.blacklistedTokens[token]
	return exists
}
