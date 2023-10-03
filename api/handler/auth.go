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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func ResetHandler(w http.ResponseWriter, r *http.Request) {

}
