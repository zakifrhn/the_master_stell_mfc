package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func VerifyPassword(hashPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	return err
}