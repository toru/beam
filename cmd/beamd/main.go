// The beamd command runs the Beam daemon.
package main

import (
	"flag"
	"log"

	"github.com/toru/beam/store"
)

type config struct {
	Store string
	Port  int
}

func buildConfig() config {
	var cfg config
	flag.StringVar(&cfg.Store, "store", "memory", "storage engine name")
	flag.IntVar(&cfg.Port, "port", 8080, "port number to listen on")
	flag.Parse()
	return cfg
}

func main() {
	cfg := buildConfig()
	_, err := store.GetStore(cfg.Store)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("looks good")
}
