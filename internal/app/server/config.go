package server

import "github.com/hzhyvinskyi/eager-beaver/internal/app/dal"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Dal      *dal.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":9092",
		LogLevel: "debug",
		Dal:	  dal.NewConfig(),
	}
}
