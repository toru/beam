// Package bookmark provides structures and functions for bookmarks.
package bookmark

import (
	"net/url"
	"time"
)

// Item represents a bookmark item.
type Item struct {
	ID          []byte    // Unique ID
	Name        string    // Name of the item
	Description string    // Description of the item
	CreatedAt   time.Time // When the item was created
	UpdatedAt   time.Time // When the item was last updated

	url    url.URL // Raw internal URL representation
	idAlgo int     // Algorithm used to generate the ID
}
