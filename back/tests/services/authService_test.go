package services

import (
	"challenges4/config"
	"challenges4/models"
	"challenges4/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Claims struct {
	UserID uint  `json:"userId"`
	Roles  uint8 `json:"roles"`
	jwt.RegisteredClaims
}

func TestGenerateJWT(t *testing.T) {
	user := models.User{
		Base:  models.Base{ID: 1},
		Roles: 1,
	}
	token, err := services.GenerateJWT(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// Validate token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	assert.Nil(t, err)
	claims, ok := parsedToken.Claims.(*Claims)
	assert.True(t, ok)
	assert.Equal(t, uint(1), claims.UserID)
	assert.Equal(t, uint8(1), claims.Roles)
}

func TestHashPassword(t *testing.T) {
	password := "test1234"
	hash, err := services.HashPassword(password)
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "test1234"
	hash, _ := services.HashPassword(password)
	assert.True(t, services.CheckPasswordHash(password, hash))
	assert.False(t, services.CheckPasswordHash("wrongpassword", hash))
}
