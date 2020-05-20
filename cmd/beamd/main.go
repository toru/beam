// The beamd command runs the Beam daemon.
package main

import (
	"log"

	"github.com/toru/beam/store"
)

func main() {
	cfg, err := buildConfig()
	if err != nil {
		log.Fatal(err)
	}
	_, err = store.GetStore(cfg.Store)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("looks good")
}
