package handler

import (
	"encoding/json"
	"net/http"
	"user-api/store"
	"user-api/util"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}

	// Get user
	user, err := store.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Check password
	if !util.CheckHashedPassword(req.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := util.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond to request
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LogoutHandler This handler has JWT Middleware; no need to check token manually
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("token")
	if value == nil {
		http.Error(w, "Token not found in context", http.StatusBadRequest)
		return
	}

	tokenStr, ok := value.(string)
	if !ok {
		http.Error(w, "Token is not of type string", http.StatusInternalServerError)
		return
	}

	// Blacklist the token
	store.AddTokenToBlacklist(tokenStr)

	// Return success response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Logged out successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
