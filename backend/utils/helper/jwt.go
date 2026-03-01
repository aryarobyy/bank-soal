package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"latih.in-be/internal/model"
)

type ClaimsModel struct {
	UserId       int    `json:"id"`
	Role         string `json:"role"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	TokenVersion int    `json:"token_version"`
	jwt.RegisteredClaims
}

func ParseExpiry(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("empty expiry string")
	}

	if strings.HasSuffix(s, "d") {
		nStr := strings.TrimSuffix(s, "d")
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return 0, err
		}
		return time.Duration(n) * 24 * time.Hour, nil
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return d, nil
}

func GenerateAccessToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expiryStr := os.Getenv("JWT_EXPIRED")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set in environment")
	}

	if expiryStr == "" {
		expiryStr = "30m"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", err
	}

	expireAt := time.Now().Add(duration)
	claims := ClaimsModel{
		UserId:       user.Id,
		Role:         string(user.Role),
		Name:         user.Name,
		Email:        userEmail,
		TokenVersion: user.TokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
			Subject:   strconv.Itoa(user.Id),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func GenerateRefreshToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	expiryStr := os.Getenv("JWT_REFRESH_EXPIRED")
	if expiryStr == "" {
		expiryStr = "7d"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", err
	}

	expireAt := time.Now().Add(duration)
	claims := ClaimsModel{
		UserId:       user.Id,
		TokenVersion: user.TokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(user.Id),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseAndValidateToken(tokenString string) (*ClaimsModel, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsModel{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsModel)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ParseTokenAllowExpired(tokenString string) (*ClaimsModel, error) {
	secret := os.Getenv("JWT_SECRET")
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	token, err := parser.ParseWithClaims(tokenString, &ClaimsModel{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsModel)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ReSignAccessToken(oldClaims *ClaimsModel) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set in environment")
	}

	expiryStr := os.Getenv("JWT_EXPIRED")
	if expiryStr == "" {
		expiryStr = "10m"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", err
	}

	oldClaims.IssuedAt = jwt.NewNumericDate(time.Now())
	oldClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, oldClaims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateRefreshToken(tokenStr string) (int, int, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &ClaimsModel{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing refresh token: %w", err)
	}

	claims, ok := token.Claims.(*ClaimsModel)
	if !ok || !token.Valid {
		return 0, 0, errors.New("invalid or expired refresh token")
	}

	if claims.Subject == "" {
		return 0, 0, errors.New("missing subject in refresh token")
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid subject format: %w", err)
	}

	return id, claims.TokenVersion, nil
}
