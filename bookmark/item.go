// Package bookmark provides structures and functions for bookmarks.
package bookmark

import (
	"net/url"
	"time"
)

const (
	AlgoSHA1 = iota
	AlgoSHA224
)

// Item represents a bookmark item.
type Item struct {
	ID          []byte    // Unique ID
	Name        string    // Name of the item
	Description string    // Description of the item
	CreatedAt   time.Time // When the item was created
	UpdatedAt   time.Time // When the item was last updated

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
	return nil
}

// URL returns a string of the item URL
func (item *Item) URL() string {
	return item.url.String()
}

// HashAlgo returns the algorithm used to generate the ID.
func (item *Item) HashAlgo() int {
	return item.hashAlgo
}
