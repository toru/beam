// The beamd command runs the Beam daemon.
package main

import (
	"log"

	"github.com/toru/beam/store"
)

func main() {
	cfg := buildConfig()
	_, err := store.GetStore(cfg.Store)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("looks good")
}
