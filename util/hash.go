package util

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plaintext password and returns its bcrypt hash.
// bcrypt is used because it is currently considered one of the most secure
// password hashing schemes, and it automatically handles the creation of salt.
//
// Parameters:
// - password: the plaintext password to be hashed.
//
// Returns:
// - the bcrypt hash of the password as a string.
// - error, if any occurred during hashing.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckHashedPassword compares a plaintext password with its bcrypt hash to
// check if they match. This is used during user login to validate the
// user-provided password against the stored hash.
//
// Parameters:
// - password: the plaintext password to check.
// - hash: the bcrypt hash against which the password needs to be checked.
//
// Returns:
// - true if the password matches the hash, false otherwise.
func CheckHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
