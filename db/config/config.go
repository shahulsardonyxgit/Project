package config
)

import (
	"log"
	"github.com/BurntSushi/toml"

// Represents database server and credentials
func (c *Config) Read() {
type Config struct {
	Server   string
	Database string
	Port int
}

// Read and parse the configuration file
	if _, err := toml.DecodeFile("configu.toml", &c); err != nil {
		log.Fatal(err)
	}
}
