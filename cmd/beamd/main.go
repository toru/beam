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
	db, _ := store.GetStore("memory")
	db.WriteBookmark(*item)
	fmt.Println(db.BookmarkCount())
}
