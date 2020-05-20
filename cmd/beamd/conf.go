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
	Store string
	Port  int
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
	flag.IntVar(&cfg.Port, "p", 8080, "port number to listen on")
	flag.Parse()
	// When a config path is present, consider the configuration file to
	// be the one and only source of configuration entries.
	if len(cfgPath) != 0 {
		return buildFromConfigFile(cfgPath)
	}
	return cfg, nil
}
