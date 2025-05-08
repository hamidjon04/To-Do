package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtKey is the secret key used to sign tokens
var JwtKey = []byte("your-secret-key") // bu kalit .env fayldan olinishi ma'qul

// Claims defines the structure for JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates and signs a JWT token for a user
func GenerateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token amal qilish muddati
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// ParseToken parses the JWT token and returns the claims
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func (c Claims) Valid() error {
	if c.ExpiresAt == nil {
		return errors.New("missing expiration claim")
	}
	if time.Now().After(c.ExpiresAt.Time) {
		return errors.New("token is expired")
	}
	return nil
}
