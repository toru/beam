// The beamd command runs the Beam daemon.
package main

import (
	"fmt"

	"github.com/toru/beam/bookmark"
)

func main() {
	item := bookmark.NewItem()
	item.Name = "Hello World"
	fmt.Println(item.Name)
}
