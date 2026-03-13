package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	Secret string
}

type UserClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (m *JWTManager) Generate(userID int64) (string, error) {

	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.Secret))
}
