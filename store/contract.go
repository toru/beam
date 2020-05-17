// Package store provides structures and functions for storage.
package store

import (
	"github.com/toru/beam/bookmark"
)

// Store is a Storage Engine interface.
type Store interface {
	Name() string                        // Name of the storage engine
	Bookmarks(limit int) []bookmark.Item // List of bookmarks
	WriteBookmark(bookmark.Item) error   // Write the given bookmark
}
