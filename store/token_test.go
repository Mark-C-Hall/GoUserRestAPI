package store

import "testing"

// TestAddTokenToBlacklist tests the function to add a token to the blacklist.
func TestAddTokenToBlacklist(t *testing.T) {
	token := "testToken"

	// Add the token to the blacklist.
	AddTokenToBlacklist(token)

	// Check if the token has been successfully added to the blacklist.
	if !IsTokenBlacklisted(token) {
		t.Errorf("Token %s was not added to the blacklist", token)
	}
}

// TestIsTokenBlacklisted tests the function to check if a token is blacklisted.
func TestIsTokenBlacklisted(t *testing.T) {
	token := "anotherTestToken"

	// Initially, the token should not be blacklisted.
	if IsTokenBlacklisted(token) {
		t.Errorf("Token %s should not be blacklisted yet", token)
	}

	// Add the token to the blacklist.
	AddTokenToBlacklist(token)

	// Now, the token should be blacklisted.
	if !IsTokenBlacklisted(token) {
		t.Errorf("Token %s should be blacklisted after adding it", token)
	}
}
