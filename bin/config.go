package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)
type Config struct {
	tap Tap
	target Target
}

type Tap struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Target struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

func (config *Config) Read(path string) {
	yamlData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("YAML Data: ", string(yamlData))

	var cfg Config

	err = yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config Full: %s\n", cfg)
	fmt.Printf("Tap Type: %s\n", cfg.tap.Type)
	fmt.Printf("Tap Name: %s\n", cfg.tap.Name)
	fmt.Printf("Tap Path: %s\n", cfg.tap.Path)
	fmt.Printf("Target Type: %s\n", cfg.target.Type)
	fmt.Printf("Target Name: %s\n", cfg.target.Name)
}