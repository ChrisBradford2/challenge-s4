package services

import (
	"challenges4/config"
	"challenges4/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	UserID uint  `json:"userId"`
	Roles  uint8 `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if the provided password matches the stored hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIDFromToken extracts the user ID from the JWT token
func GetUserIDFromToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["userId"].(float64))
		return userID, nil
	} else {
		return 0, err
	}
}
