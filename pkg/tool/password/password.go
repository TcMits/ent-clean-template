package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword validates a plain password against the model's password.
func ValidatePassword(passwordHash, password string) error {
	bytePassword := []byte(password)
	bytePasswordHash := []byte(passwordHash)

	// comparing the password with the hash
	return bcrypt.CompareHashAndPassword(bytePasswordHash, bytePassword)
}

// SetPassword sets cryptographically secure string to `model.Password`.
func GetHashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("The provided plain password is empty")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
