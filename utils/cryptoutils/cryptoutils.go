package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword hashes a password and returns the
// hash
func HashPassword(password string) string {

	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
