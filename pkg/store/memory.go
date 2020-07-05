package store

import (
	"sync"
	"time"

	"github.com/toru/beam/pkg/auth"
	"github.com/toru/beam/pkg/bookmark"
)

// MemoryStore is an in-memory storage engine. Consider it as a reference
// implementation. MemoryStore is unsuited for production.
type MemoryStore struct {
	mux       sync.RWMutex
	authKeys  map[string]auth.Key
	bookmarks map[string]bookmark.Item
}

// NewMemoryStore returns a pointer to a new MemoryStore
func NewMemoryStore(_ string) *MemoryStore {
	ret := &MemoryStore{}
	ret.bookmarks = make(map[string]bookmark.Item)
	return ret
}

// Name implements the Store interface
func (s MemoryStore) Name() string {
	return "Memory Store"
}

// Bookmarks implements the Store interface
func (s MemoryStore) Bookmarks(limit int) []bookmark.Item {
	s.mux.RLock()
	defer s.mux.RUnlock()

	// TODO(toru): Tweak the initial allocation
	bookmarks := make([]bookmark.Item, 0, len(s.bookmarks))
	for _, b := range s.bookmarks {
		bookmarks = append(bookmarks, b)
	}
	return bookmarks
}

// BookmarkCount implements the Store interface
func (s MemoryStore) BookmarkCount() int {
	return len(s.bookmarks)
}

// GetBookmark implements the Store interface
func (s MemoryStore) GetBookmark(id string) (bookmark.Item, bool) {
	s.mux.RLock()
	item, ok := s.bookmarks[id]
	s.mux.RUnlock()
	return item, ok
}

// WriteBookmark implements the Store interface
func (s MemoryStore) WriteBookmark(item *bookmark.Item) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	now := time.Now().UTC()
	if item.CreatedAt.IsZero() {
		item.CreatedAt = now
	}
	item.UpdatedAt = now
	s.bookmarks[item.HexID()] = item.Dup()
	return nil
}

// WriteAuthKey implements the Store interface
func (s MemoryStore) GetAuthKey(token string) (auth.Key, bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	key, ok := s.authKeys[token]
	return key, ok
}

// WriteAuthKey implements the Store interface
func (s MemoryStore) WriteAuthKey(key *auth.Key) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	now := time.Now().UTC()
	if key.CreatedAt.IsZero() {
		key.CreatedAt = now
	}
	key.UpdatedAt = now
	s.authKeys[key.HexToken()] = key.Dup()
	return nil
}
