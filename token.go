package test

import (
	"Nikcase/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "grkjk#4#%35FSFJLja#4353KSFjH"
	lifetime   = 3 * time.Hour
)

func CreateToken(login string) (token string, err error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: login,
	})
	token, err = tok.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
