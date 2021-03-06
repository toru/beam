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

func addrLabel(src string) string {
	if len(src) == 0 {
		return "*"
	}
	return src
}

func main() {
	cfg, err := buildConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := store.GetStore(cfg.Store, cfg.DataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("looks good")

	lsnr, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	crash := make(chan error, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("starting beamd... addr:%s, port:%d, tls:%t",
		addrLabel(cfg.Listen), cfg.Port, cfg.canServeTLS())
	go func() {
		app := NewBeamApp(cfg, db)
		if cfg.canServeTLS() {
			crash <- http.ServeTLS(lsnr, app, cfg.CertPath, cfg.KeyPath)
		} else {
			crash <- http.Serve(lsnr, app)
		}
	}()

	select {
	case err := <-crash:
		log.Printf("web server crashed: %v", err)
	case sig := <-sigs:
		log.Printf("received %s, quitting", sig.String())
	}
}
