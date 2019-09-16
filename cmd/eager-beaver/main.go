package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/hzhyvinskyi/eager-beaver/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
