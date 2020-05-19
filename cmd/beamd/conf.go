package main

import (
	"flag"
)

type config struct {
	Store string
	Port  int
}

func buildConfig() config {
	var cfg config
	flag.StringVar(&cfg.Store, "store", "memory", "storage engine name")
	flag.IntVar(&cfg.Port, "p", 8080, "port number to listen on")
	flag.Parse()
	return cfg
}
