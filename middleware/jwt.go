package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"user-api/store"
	"user-api/util"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := extractTokenFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if store.IsTokenBlacklisted(tokenStr) {
			http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
			return
		}

		claims, err := util.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		ctx = context.WithValue(ctx, "token", tokenStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func extractTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("malformed authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		return "", errors.New("token not provided")
	}

	return tokenStr, nil
}
