package util

import (
	"testing"
)

// TestHashAndCheckPassword tests the utility functions for hashing and checking hashed passwords.
// This test uses a sample password, hashes it, and then checks the hashed password.
func TestHashAndCheckPassword(t *testing.T) {
	password := "testPassword123"

	// Hash the sample password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatalf("Expected a hashed password, got an empty string")
	}

	// Check the hashed password with the original password
	isValid := CheckHashedPassword(password, hashedPassword)
	if !isValid {
		t.Fatal("Failed to validate the hashed password")
	}
}

// TestCheckWrongPassword tests the CheckHashedPassword function with a wrong password.
// This test is expected to fail password validation.
func TestCheckWrongPassword(t *testing.T) {
	password := "testPassword123"
	wrongPassword := "wrongPassword123"

	// Hash the sample password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Attempt to check the hashed password with a wrong password
	isValid := CheckHashedPassword(wrongPassword, hashedPassword)
	if isValid {
		t.Fatal("Expected password validation to fail, but it passed")
	}
}
