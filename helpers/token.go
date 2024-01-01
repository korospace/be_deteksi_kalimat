package helpers

import (
	"be_deteksi_kalimat/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("03102000")

type TokenInfo struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := TokenInfo{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ValidateToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenInfo{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	// token is valid
	claims, ok := token.Claims.(*TokenInfo)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	// token is expired
	now := time.Now()
	expirationTime := time.Unix(claims.ExpiresAt.Time.Unix(), 0)
	if now.After(expirationTime) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
