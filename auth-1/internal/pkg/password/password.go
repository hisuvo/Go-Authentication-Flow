package password

import "golang.org/x/crypto/bcrypt"

const cost = 12

// Hash creates a secure password hash.
func Hash(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword),cost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare verifies the password.
func Compare(hash, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
}