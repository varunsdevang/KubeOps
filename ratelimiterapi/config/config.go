package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}

func loadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(data, &Config{}) // Unmarshal to a temporary config struct
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	config = &Config{}                 // assign to global config.
	err = yaml.Unmarshal(data, config) // Unmarshal data into global config.
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return nil
}

func GetConfig(path string) (*Config, error) {
	var err error
	once.Do(func() {
		err = loadConfig(path) // Load config only once
	})

	if err != nil {
		return nil, err
	}
	return config, nil
}
