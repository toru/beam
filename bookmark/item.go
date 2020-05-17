// Package bookmark provides structures and functions for bookmarks.
package bookmark

import (
	"crypto/sha1"
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

// SetURL sets the given URL string to the item
func (item *Item) SetURL(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	item.url = *u
	item.id = hash(item.hashAlgo, urlStr)
	return nil
}

// ID returns the unique item ID
func (item *Item) ID() []byte {
	return item.id
}

// URL returns a string of the item URL
func (item *Item) URL() string {
	return item.url.String()
}

// HashAlgo returns the algorithm used to generate the ID.
func (item *Item) HashAlgo() int {
	return item.hashAlgo
}

func hash(algo int, urlStr string) []byte {
	if algo != AlgoSHA1 {
		return nil
	}
	digest := sha1.Sum([]byte(urlStr))
	return digest[:]
}
