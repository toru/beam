// The beamd command runs the Beam daemon.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	log.Println("looks good")

	lsnr, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	crash := make(chan error, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("starting beamd... port:%d, tls:%t", cfg.Port, cfg.canServeTLS())
	go func() {
		if cfg.canServeTLS() {
			crash <- http.ServeTLS(lsnr, nil, cfg.CertPath, cfg.KeyPath)
		} else {
			crash <- http.Serve(lsnr, nil)
		}
	}()

	select {
	case err := <-crash:
		log.Printf("web server crashed: %v", err)
	case sig := <-sigs:
		log.Printf("received %s, quitting", sig.String())
	}
}
