// Package config provides for parse config file.

package config

type Config struct {
	globalContent   map[string]string
	sectionContents map[string]map[string]string
	sections        []string
}

func NewConfig() *Config {
	return &Config{
		globalContent:   make(map[string]string),
		sectionContents: make(map[string]map[string]string),
	}
}
