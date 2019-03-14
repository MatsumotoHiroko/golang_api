package config

import (
	"log"
	"flag"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	var toml_path string
	if flag.Lookup("test.v") != nil {
        toml_path = "config_test.toml"
	} else {
        toml_path = "config.toml"
	}
	if _, err := toml.DecodeFile(toml_path, &c); err != nil {
		log.Fatal(err)
	}
}