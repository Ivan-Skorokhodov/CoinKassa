package hash

import (
	"CoinKassa/internal/models"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func hashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 16*1024, 2, 16)
	return base64.StdEncoding.EncodeToString(hash)
}

func createSalt() ([]byte, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func CreateUID(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func HashPasswordAndCreateSalt(password string, store *models.Store) error {
	salt, err := createSalt()
	if err != nil {
		return err
	}

	hashedPassword := hashPassword(password, salt)
	store.PasswordHash = hashedPassword
	store.Salt = base64.RawStdEncoding.EncodeToString(salt)

	return nil
}

func CheckPassword(password, passwordFromDB string, salt []byte) bool {
	hashedInput := hashPassword(password, salt)
	return subtle.ConstantTimeCompare([]byte(hashedInput), []byte(passwordFromDB)) == 1
}
