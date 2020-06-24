package store

import (
	"sort"
	"testing"

	"github.com/toru/beam/pkg/bookmark"
)

func TestName(t *testing.T) {
	db := NewMemoryStore("/tmp/beam")
	got := db.Name()
	want := "Memory Store"
	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}

func TestBookmarks(t *testing.T) {
	db := NewMemoryStore("/tmp/beam")
	i1 := bookmark.NewItem()
	i1.SetURL("https://torumk.com/")
	i2 := bookmark.NewItem()
	i2.SetURL("https://ep.torumk.com/")
	db.WriteBookmark(i1)
	db.WriteBookmark(i2)

	// Hack until we implement comparators.
	keys := make([]string, 0, 2)
	for _, k := range db.Bookmarks(0) {
		keys = append(keys, k.URL())
	}
	sort.Strings(keys)
	want := "https://ep.torumk.com/"
	if keys[0] != want {
		t.Errorf("Got: %s, Want: %s", keys[0], want)
	}
	want = "https://torumk.com/"
	if keys[1] != want {
		t.Errorf("Got: %s, Want: %s", keys[1], want)
	}
}

func TestBookmarkCount(t *testing.T) {
	db := NewMemoryStore("/tmp/beam")
	i1 := bookmark.NewItem()
	i1.SetURL("https://torumk.com/")
	i2 := bookmark.NewItem()
	i2.SetURL("https://ep.torumk.com/")
	db.WriteBookmark(i1)
	db.WriteBookmark(i2)
	got := db.BookmarkCount()
	want := 2
	if got != want {
		t.Errorf("Got: %d, Want: %d", got, want)
	}
}

func TestGetBookmark(t *testing.T) {
	db := NewMemoryStore("/tmp/beam")
	item := bookmark.NewItem()
	item.SetURL("https://torumk.com/")
	if err := db.WriteBookmark(item); err != nil {
		t.Error(err)
	}
	_, ok := db.GetBookmark("bogus")
	if ok {
		t.Errorf("Got: %t, Want: false", ok)
	}
	_, ok = db.GetBookmark(item.HexID())
	if !ok {
		t.Errorf("Got: %t, Want: true", ok)
	}
}

func TestWriteBookmark(t *testing.T) {
	db := NewMemoryStore("/tmp/beam")
	item := bookmark.NewItem()
	item.SetURL("https://torumk.com/")
	if err := db.WriteBookmark(item); err != nil {
		t.Error(err)
	}
}
