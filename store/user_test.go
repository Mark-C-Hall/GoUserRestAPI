package store

import (
	"testing"
	"user-api/util"
)

func TestAddAndGetUser(t *testing.T) {
	tests := []struct {
		user     User
		expected string
	}{
		{User{Username: "TestUser1", Email: "test1@email.com", Password: "password1"}, "TestUser1"},
		{User{Username: "TestUser2", Email: "test2@email.com", Password: "password2"}, "TestUser2"},
	}

	for _, tt := range tests {
		err := CreateUser(&tt.user)
		if err != nil {
			t.Fatalf("Failed to create new user: %v", err)
		}

		retrievedUser, err := GetUserByUsername(tt.user.Username)
		if err != nil {
			t.Fatalf("Failed to retrieve user: %v", err)
		}

		if tt.user.Username != retrievedUser.Username {
			t.Fatalf("Expected %s, but got %s", tt.expected, retrievedUser.Username)
		}
	}
}

func TestNonExistentUser(t *testing.T) {
	_, err := GetUserByUsername("Nobody")
	if err == nil {
		t.Fatalf("Expected error for non-existent user, but got none")
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup: Create a user to be updated
	user := User{
		Username: "UpdateTestUser",
		Email:    "UpdateTest@email.com",
		Password: "initialPassword",
	}
	err := CreateUser(&user)
	if err != nil {
		t.Fatalf("Failed to create user for update: %v", err)
	}

	// Update the user's details
	updatedDetails := User{
		Username: user.Username, // Username remains the same, as it's the key
		Email:    "UpdatedTest@email.com",
		Password: "updatedPassword",
	}
	err = UpdateUser(&updatedDetails)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Retrieve the user and validate the updated details
	retrievedUser, err := GetUserByUsername(user.Username)
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}

	if retrievedUser.Email != updatedDetails.Email {
		t.Fatalf("Expected updated email %s, got %s", updatedDetails.Email, retrievedUser.Email)
	}

	// You'd probably use a function to check the hashed password matches the plain-text password here
	if !util.CheckHashedPassword(updatedDetails.Password, retrievedUser.Password) {
		t.Fatalf("Password was not updated correctly")
	}
}

func TestDeleteUser(t *testing.T) {
	// Setup: Create a user to be deleted
	user := User{
		Username: "DeleteTestUser",
		Email:    "DeleteTest@email.com",
		Password: "deleteMePassword",
	}
	err := CreateUser(&user)
	if err != nil {
		t.Fatalf("Failed to create user for deletion: %v", err)
	}

	// Delete the user
	err = DeleteUserByUsername(user.Username)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Try to retrieve the user; expect an error since the user should no longer exist
	_, err = GetUserByUsername(user.Username)
	if err == nil {
		t.Fatal("Expected error retrieving deleted user, but got none")
	}
}
