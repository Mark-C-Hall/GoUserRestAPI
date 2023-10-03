package util

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "testPassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatalf("Expected a hashed password, got an empty string")
	}

	isValid := CheckHashedPassword(password, hashedPassword)
	if !isValid {
		t.Fatal("Failed to validate the hashed password")
	}
}

func TestCheckWrongPassword(t *testing.T) {
	password := "testPassword123"
	wrongPassword := "wrongPassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	isValid := CheckHashedPassword(wrongPassword, hashedPassword)
	if isValid {
		t.Fatal("Expected password validation to fail, but it passed")
	}
}
