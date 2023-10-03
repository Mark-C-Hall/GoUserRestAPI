package handler

import (
	"encoding/json"
	"net/http"
	"user-api/store"
	"user-api/util"
)

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

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

	// Generate the token
	token, err := util.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond to request with the generated token
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"message": "User registered successfully",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ProfileHandler This handler has JWT Middleware; no need to check token manually
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*util.Claims)
	user, err := store.GetUserByUsername(claims.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user profile data; mask sensitive data
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateUserHandler This handler has JWT Middleware; no need to check token manually
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*util.Claims)

	var updatedUser store.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Make sure the updated user matches the authenticated user
	updatedUser.Username = claims.Username

	// Update user in the store
	err = store.UpdateUser(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteUserHandler This handler has JWT Middleware; no need to check token manually
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*util.Claims)

	err := store.DeleteUserByUsername(claims.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
