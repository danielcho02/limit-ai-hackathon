package auth

import (
	"errors"
	"main/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserID int `json:"user_id"`
	Role	domain.RoleType `json:"role"`
	jwt.RegisteredClaims
}

const (
	DefaultIssuer = "dasom.io"
)

// Token expires at expirationTime
func GenerateJWT(userID int, role domain.RoleType, secretKey []byte, duration time.Duration) (string, error) {
	claims := &customClaims{
		UserID: userID,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    DefaultIssuer,
		},
	}
	// hashing algorithm : HS256(HMAC + SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token == nil {
		return "", jwt.ErrTokenMalformed
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokens(userID int, role domain.RoleType, secretKey []byte) (accessToken string, refreshToken string, err error) {
	// Access token expires in 2 hours
	accessToken, err = GenerateJWT(userID, role, secretKey, 2*time.Hour)
	if err != nil {
		return "", "", err
	}

	// Refresh token expires in 14 days
	refreshToken, err = GenerateJWT(userID, role, secretKey, 14*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateJWT(tokenString string, jwtSecret []byte) (*customClaims, error) {
	claims := &customClaims{}

	// Parse and validate
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check issuer (defensive)
	if claims.Issuer != DefaultIssuer {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}