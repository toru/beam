// Package store provides structures and functions for storage.
package store

import (
	"errors"

	"github.com/toru/beam/bookmark"
)

// Store is a Storage Engine interface.
type Store interface {
	Name() string                        // Name of the storage engine
	Bookmarks(limit int) []bookmark.Item // List of bookmarks
	BookmarkCount() int                  // Number of bookmarks
	WriteBookmark(bookmark.Item) error   // Write the given bookmark
}

// GetStore returns the specified Storage Engine.
func GetStore(name string) (Store, error) {
	if name != "memory" {
		return nil, errors.New("unknown storage engine")
	}
	return NewMemoryStore(), nil
}
