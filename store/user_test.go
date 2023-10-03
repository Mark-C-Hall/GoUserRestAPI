package store

import (
	"testing"
	"user-api/util"
)

// TestAddAndGetUser tests the functionality of adding a user to the store and then retrieving them.
// It uses table-driven testing to iterate over a slice of test cases, creating a user for each and
// ensuring the added user can be retrieved by username.
func TestAddAndGetUser(t *testing.T) {
	tests := []struct {
		user     User
		expected string
	}{
		{User{Username: "TestUser1", Email: "test1@email.com", Password: "password1"}, "TestUser1"},
		{User{Username: "TestUser2", Email: "test2@email.com", Password: "password2"}, "TestUser2"},
	}

	for _, tt := range tests {
		// Create the user
		err := CreateUser(&tt.user)
		if err != nil {
			t.Fatalf("Failed to create new user: %v", err)
		}

		// Retrieve and validate the created user
		retrievedUser, err := GetUserByUsername(tt.user.Username)
		if err != nil {
			t.Fatalf("Failed to retrieve user: %v", err)
		}
		if tt.user.Username != retrievedUser.Username {
			t.Fatalf("Expected %s, but got %s", tt.expected, retrievedUser.Username)
		}
	}
}

// TestNonExistentUser tests the behavior of trying to retrieve a user that doesn't exist.
// The expected behavior is that an error should be returned.
func TestNonExistentUser(t *testing.T) {
	_, err := GetUserByUsername("Nobody")
	if err == nil {
		t.Fatalf("Expected error for non-existent user, but got none")
	}
}

// TestUpdateUser tests the user update functionality.
// It begins by creating a user, updates the user's email and password, and then validates
// that the updates were applied correctly in the store.
func TestUpdateUser(t *testing.T) {
	// Create a user for testing the update
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
		Username: user.Username, // Using the same username since it's our retrieval key
		Email:    "UpdatedTest@email.com",
		Password: "updatedPassword",
	}
	err = UpdateUser(&updatedDetails)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Retrieve and validate the updated user details
	retrievedUser, err := GetUserByUsername(user.Username)
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}
	if retrievedUser.Email != updatedDetails.Email {
		t.Fatalf("Expected updated email %s, got %s", updatedDetails.Email, retrievedUser.Email)
	}
	if !util.CheckHashedPassword(updatedDetails.Password, retrievedUser.Password) {
		t.Fatalf("Password was not updated correctly")
	}
}

// TestDeleteUser tests the user deletion functionality.
// It begins by creating a user, deletes it, and then ensures that the user no longer exists in the store.
func TestDeleteUser(t *testing.T) {
	// Create a user for testing the delete functionality
	user := User{
		Username: "DeleteTestUser",
		Email:    "DeleteTest@email.com",
		Password: "deleteMePassword",
	}
	err := CreateUser(&user)
	if err != nil {
		t.Fatalf("Failed to create user for deletion: %v", err)
	}

	// Delete the created user
	err = DeleteUserByUsername(user.Username)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Ensure that the deleted user can't be retrieved
	_, err = GetUserByUsername(user.Username)
	if err == nil {
		t.Fatal("Expected error retrieving deleted user, but got none")
	}
}
