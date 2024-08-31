package security

import (
	"RD-Clone-NAPI/internal/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTokenWithExp generates a JWT with an expiration of 1 hour (exp time comes from the config).
func GenerateTokenWithExp(email string) (string, time.Time, error) {
	jwtConfig := config.Load().JWT

	currentTime := time.Now().Local()
	expirationDate := currentTime.Add(time.Second * time.Duration(jwtConfig.Expiration))

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationDate),
		IssuedAt:  jwt.NewNumericDate(currentTime),
		Issuer:    "GO-Reddit-CL",
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(jwtConfig.Key))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("could not generate JWT %w please try again", err)
	}

	return signedToken, expirationDate, nil
}
