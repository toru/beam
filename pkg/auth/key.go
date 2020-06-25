// Package auth provides code for API authentication and authorization.
package auth

// Key represents an API key.
type Key struct {
	Name       string    // Name of the key
	Token      []byte    // Bearer token
	LastUsedAt time.Time // When the key was last used
	CreatedAt  time.Time // When the key was created
	UpdatedAt  time.Time // When the key was last updated
}
