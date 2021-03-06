// Package bookmark provides structures and functions for bookmarks.
package bookmark

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"time"
)

const (
	AlgoSHA1 = iota
	AlgoSHA224
)

// Item represents a bookmark item.
type Item struct {
	Name        string    // Name of the item
	Description string    // Description of the item
	CreatedAt   time.Time // When the item was created
	UpdatedAt   time.Time // When the item was last updated

	id       []byte  // Unique ID
	url      url.URL // Raw internal URL representation
	hashAlgo int     // Algorithm used to generate the ID
}

// NewItem returns a pointer to a new Item.
func NewItem() *Item {
	return &Item{hashAlgo: AlgoSHA1}
}

// SetURL sets the given URL string to the item.
func (item *Item) SetURL(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	item.url = *u
	item.id = hash(item.hashAlgo, urlStr)
	return nil
}

// ID returns the unique item ID.
func (item *Item) ID() []byte {
	return item.id
}

// HexID returns the hexadecimal string representation of the item ID.
func (item *Item) HexID() string {
	return hex.EncodeToString(item.id)
}

// URL returns a string of the item URL.
func (item *Item) URL() string {
	return item.url.String()
}

// HashAlgo returns the algorithm used to generate the ID.
func (item *Item) HashAlgo() int {
	return item.hashAlgo
}

// Dup returns a copy of the given Item pointer value.
func (item *Item) Dup() Item {
	dup := *item
	dup.id = make([]byte, len(item.id))
	copy(dup.id, item.id)
	return dup
}

// MarshalJSON implements the json.Marsharler interface.
func (item Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        string    `json:"id"`
		URL       string    `json:"url"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        item.HexID(),
		URL:       item.URL(),
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	})
}

func hash(algo int, urlStr string) []byte {
	if algo != AlgoSHA1 {
		return nil
	}
	digest := sha1.Sum([]byte(urlStr))
	return digest[:]
}
