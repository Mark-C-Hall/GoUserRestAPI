package util

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	username := "TestUser"
	tokenStr, err := GenerateToken(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %s, but got %s", username, claims.Username)
	}

	// Test with an invalid token
	_, err = ValidateToken(tokenStr + "invalid")
	if err == nil {
		t.Fatalf("Expected error for invalid token, but got none")
	}
}

func TestTokenExpiry(t *testing.T) {
	username := "TestUser"

	// Create a token that expires immediately
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Second)),
		},
	})

	tokenStr, err := expiredToken.SignedString(JWTKey)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Fatalf("Expected error for expired token, but got none")
	}
}
