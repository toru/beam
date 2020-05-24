// The beamd command runs the Beam daemon.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/toru/beam/pkg/store"
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
	log.Printf("TLS: %t\n", cfg.canServeTLS())
	log.Println("looks good")

	lsnr, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	// TODO(toru): Run the web app in a dedicated goroutine.
	if cfg.canServeTLS() {
		log.Fatal(http.ServeTLS(lsnr, nil, cfg.CertPath, cfg.KeyPath))
	} else {
		log.Fatal(http.Serve(lsnr, nil))
	}
}
