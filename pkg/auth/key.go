// Package auth provides code for API authentication and authorization.
package auth

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"
)

// Key represents an API key.
type Key struct {
	Name       string    // Name of the key
	Token      []byte    // Bearer token
	LastUsedAt time.Time // When the key was last used
	CreatedAt  time.Time // When the key was created
	UpdatedAt  time.Time // When the key was last updated
}

// NewKey returns a pointer to a new Key.
func NewKey() *Key {
	return &Key{Token: genToken()}
}

// HexToken returns the hexadecimal string representation of the token.
func (key *Key) HexToken() string {
	return hex.EncodeToString(key.Token)
}

// Dup returns a copy of the given Key pointer value.
func (key *Key) Dup() Key {
	dup := *key
	dup.Token = make([]byte, len(key.Token))
	copy(dup.Token, key.Token)
	return dup
}

func genToken() []byte {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	return b
}
