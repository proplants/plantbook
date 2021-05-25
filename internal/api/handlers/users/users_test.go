// Package users contains handler implementations of the ./internal/api/restapi/operations/user
package users

import (
	"crypto/rand"
	"testing"
)

var (
	salt          []byte = make([]byte, SaltLen)
	plainPassword string = "password"
)

func TestHashPass(t *testing.T) {
	tests := []string{"pass1", "love", "$Tr0nGP@$$w0rd123", "____PASS____", "empty--", "        ", "", "password"}
	for _, plainPass := range tests {
		t.Run(plainPass, func(t *testing.T) {
			rand.Read(salt)
			passBts := HashPass(salt, plainPass)
			if !CheckPass(passBts, plainPass) {
				t.Error("unexpected not equal hashs passwords")
			}
			t.Logf("for pass: %s hash is %s\n", plainPass, string(passBts))
		})
	}
}
