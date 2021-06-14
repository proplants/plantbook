// Package randutil contains functions generate random sequences of runes or bytes
package randutil

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// RandStringRunes generates random string specified length.
// nolint:gosec
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandBytesHex generates random string specified length in hex format,
// rand is pseudo-random number generators,
// for security-sensitive work, see the RandCryptBytesHex with crypto/rand package.
func RandBytesHex(n int) string {
	return fmt.Sprintf("%x", RandBytes(n))
}

// RandCryptBytesHex generates random string specified length in hex format,
// suitable for security-sensitive work.
func RandCryptBytesHex(n int) string {
	return fmt.Sprintf("%x", RandCryptBytes(n))
}

// RandBytes generates random []byte specified length,
// rand is pseudo-random number generators,
// for security-sensitive work, see the RandCryptBytes with crypto/rand package.
// nolint:gosec
func RandBytes(n int) []byte {
	res := make([]byte, n)
	rand.Read(res)
	return res
}

// RandCryptBytes generates random []byte specified length,
// suitable for security-sensitive work.
// nolint:errcheck
func RandCryptBytes(n int) []byte {
	res := make([]byte, n)
	crand.Read(res)
	return res
}
