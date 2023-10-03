package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-api/config"
)

// JWTKey is the secret key used to sign and validate JWTs.
// It's sourced from configuration.
var JWTKey = []byte(config.C.JWTSecret)

// Claims defines the structure for JWT claims for the API.
// It embeds jwt.RegisteredClaims to include standard claims.
type Claims struct {
	Username string
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a given username.
// The token will expire 24 hours from the time of generation.
//
// Parameters:
// - username: the name of the user for whom the token is being generated.
//
// Returns:
// - a JWT as a string.
// - error, if any occurred during token generation.
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

// ValidateToken verifies the JWT's signature and claims (like expiration).
// If valid, it returns the decoded claims.
//
// Parameters:
// - tokenStr: the JWT as a string.
//
// Returns:
// - a pointer to the Claims structure if the token is valid.
// - error if the token is invalid or if there's any other error.
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
