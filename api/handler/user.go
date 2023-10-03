package handler

import (
	"encoding/json"
	"net/http"
	"user-api/store"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user store.User

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}

	// Store user in data store
	err = store.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond to request
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

}
