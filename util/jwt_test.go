package util

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

// TestGenerateAndValidateToken tests the token generation and validation process.
// It first generates a token for a sample username, then validates it to ensure the content is correct.
// Finally, it tests an invalid token to make sure validation catches the problem.
func TestGenerateAndValidateToken(t *testing.T) {
	username := "TestUser"

	// Generate a token for the test username
	tokenStr, err := GenerateToken(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the generated token and check the claims
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %s, but got %s", username, claims.Username)
	}

	// Test token validation with an invalid token string
	_, err = ValidateToken(tokenStr + "invalid")
	if err == nil {
		t.Fatalf("Expected error for invalid token, but got none")
	}
}

// TestTokenExpiry tests the token validation process with an expired token.
// It creates a token set to expire immediately, then validates it.
// The validation is expected to fail since the token is expired.
func TestTokenExpiry(t *testing.T) {
	username := "TestUser"

	// Generate an expired token for the test username
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

	// Validate the expired token and expect an error
	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Fatalf("Expected error for expired token, but got none")
	}
}
