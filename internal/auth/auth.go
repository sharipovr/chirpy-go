package auth

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	// : Hash the password using the argon2id.CreateHash function.
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	// : Use the argon2id.ComparePasswordAndHash function to compare the password that the user entered in the HTTP request with the password that is stored in the database

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}
