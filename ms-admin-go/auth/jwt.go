package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomClaims incluye user_id y role para uso posterior
type CustomClaims struct {
	UserID uuid.UUID `json:"userId"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	keyBytes := []byte(secret)

	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}

	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return keyBytes, nil
	})

	if err != nil {
		fmt.Println("ParseWithClaims error:", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Println("Token invalid")
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		fmt.Println("Invalid claims type")
		return nil, errors.New("invalid claims")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
