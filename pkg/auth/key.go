// Package auth provides code for API authentication and authorization.
package auth

import (
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
	return &Key{}
}

// Dup returns a copy of the given Key pointer value.
func (key *Key) Dup() Key {
	dup := *key
	dup.Token = make([]byte, len(key.Token))
	copy(dup.Token, key.Token)
	return dup
}
