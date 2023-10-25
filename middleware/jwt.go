package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"user-api/store"
	"user-api/util"
)

var isTokenBlacklisted = store.IsTokenBlacklisted
var validateToken = util.ValidateToken

// JWTMiddleware ensures that the provided JWT in the request header is valid,
// not blacklisted, and puts its contents (claims and token) into the request's context.
// If the token is not valid, it will respond with a 401 Unauthorized status.
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header
		tokenStr, err := extractTokenFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if the token is blacklisted
		if isTokenBlacklisted(tokenStr) {
			http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
			return
		}

		// Validate the token to get its claims
		claims, err := validateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims and token to the request context
		ctx := context.WithValue(r.Context(), "claims", claims)
		ctx = context.WithValue(ctx, "token", tokenStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// extractTokenFromRequest retrieves the JWT token from the "Authorization" header.
// It expects the header to have the format "Bearer <token>".
// Returns an error if the header is missing, malformed, or does not contain a token.
func extractTokenFromRequest(r *http.Request) (string, error) {
	// Get the "Authorization" header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	// Check if the header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("malformed authorization header")
	}

	// Extract the token from the header
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		return "", errors.New("token not provided")
	}

	return tokenStr, nil
}
