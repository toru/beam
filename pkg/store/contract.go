// Package store provides structures and functions for storage.
package store

import (
	"fmt"

	"github.com/toru/beam/pkg/bookmark"
)

// Store is a Storage Engine interface.
type Store interface {
	Name() string                                // Name of the storage engine
	Bookmarks(limit int) []bookmark.Item         // List of bookmarks
	BookmarkCount() int                          // Number of bookmarks
	GetBookmark(id string) (bookmark.Item, bool) // Get a bookmark by its unique ID
	WriteBookmark(*bookmark.Item) error          // Write the given bookmark
}

// GetStore returns the specified Storage Engine.
func GetStore(name, dataPath string) (Store, error) {
	if name != "memory" {
		return nil, fmt.Errorf("unknown storage engine: %s", name)
	}
	return NewMemoryStore(dataPath), nil
}
