package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
}

func (s BcryptService) Hash(data string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s BcryptService) Compare(data string, encrypted string) error {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(data))
}
