// The beamd command runs the Beam daemon.
package main

import (
	"fmt"

	"github.com/toru/beam/bookmark"
	"github.com/toru/beam/store"
)

func main() {
	item := bookmark.NewItem()
	item.SetURL("https://torumk.com")
	fmt.Println(item.ID())
	fmt.Println(item.HexID())
	fmt.Println(item.URL())

	db := store.NewMemoryStore()
	db.WriteBookmark(*item)
	fmt.Println(db.Name())
	fmt.Println(db.BookmarkCount())

	for _, b := range db.Bookmarks(0) {
		fmt.Println(b.HexID())
	}
}
