// Package users contains handler implementations of the ./internal/api/restapi/operations/user
package users

import (
	"bytes"
	"context"
	"crypto/rand"

	"github.com/kaatinga/plantbook/internal/api/models"

	"golang.org/x/crypto/argon2"
)

const (
	SaltLen int = 16
)

// RepoInterface users repository behavior
type RepoInterface interface {
	StoreUser(ctx context.Context, user *models.User, passwordHash []byte) (*models.User, error)
	FindUserByLogin(ctx context.Context, login string) (*models.User, []byte, error)
}

// HashPass calculate password hash with salt
func HashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

// CheckPass compare hash and string password
func CheckPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, SaltLen)
	copy(salt, passHash)
	userHash := HashPass(salt, plainPassword)
	return bytes.Equal(userHash, passHash)
}

// helpers

// nolint:errcheck
func MakeSalt(n int) []byte {
	salt := make([]byte, n)
	rand.Read(salt)
	return salt
}
