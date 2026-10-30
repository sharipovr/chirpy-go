package auth

import (
	"time"

	"github.com/google/uuid"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

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

// Using JWT token for authenticatio
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn).UTC()),
		Subject:   userID.String(),
	})

	tokenSigned, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {

	claims := jwt.RegisteredClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil })
	if err != nil {
		emptyUUID := uuid.UUID{}
		return emptyUUID, err
	}

	id := claims.Subject
	userUUID, err := uuid.Parse(id)
	if err != nil {
		emptyUUID := uuid.UUID{}
		return emptyUUID, err
	}
	return userUUID, nil
}
