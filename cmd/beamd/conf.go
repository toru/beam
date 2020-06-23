package main

import (
	"flag"

	"github.com/pelletier/go-toml"
)

const (
	defaultPort  = 8080
	defaultStore = "memory"
)

type config struct {
	Store    string // Storage Engine name
	Listen   string // TCP network address to listen on
	Port     int    // Port to listen on
	CertPath string `toml:"cert_path"` // Path to the TLS cert
	KeyPath  string `toml:"key_path"`  // Path to the private key
	DataPath string `toml:"data_path"` // Path to the data directory
}

func fillMissingWithDefault(cfg *config) {
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	if len(cfg.Store) == 0 {
		cfg.Store = defaultStore
	}
}

func buildFromConfigFile(path string) (config, error) {
	var cfg config
	cfgTree, err := toml.LoadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := cfgTree.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	fillMissingWithDefault(&cfg)
	return cfg, nil
}

func buildConfig() (config, error) {
	var cfg config
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "path to the configuration file")
	flag.StringVar(&cfg.Store, "store", "memory", "storage engine name")
	flag.StringVar(&cfg.CertPath, "cert", "", "path to the tls certificate")
	flag.StringVar(&cfg.KeyPath, "key", "", "path to the private key")
	flag.StringVar(&cfg.Listen, "l", "", "tcp network address to listen on")
	flag.IntVar(&cfg.Port, "p", 8080, "port number to listen on")
	flag.Parse()
	// When a config path is present, consider the configuration file to
	// be the one and only source of configuration entries.
	if len(cfgPath) != 0 {
		return buildFromConfigFile(cfgPath)
	}
	return cfg, nil
}

func (c config) canServeTLS() bool {
	return len(c.CertPath) != 0 && len(c.KeyPath) != 0
}
