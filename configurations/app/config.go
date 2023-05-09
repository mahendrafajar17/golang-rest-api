package app

import (
	"fmt"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Path struct {
		Encrypt string `yaml:"encrypt"`
	} `yaml:"path"`
}

func (c Config) BASEURL() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
