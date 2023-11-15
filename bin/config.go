package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"encoding/json"
	"os"
	"log"
)
type Config struct {
	tap Tap
	target Target
}

type Tap struct {
	Type string
	Name string
	Path string
}

type Target struct {
	Type string
	Name string
}

func (config *Config) ReadYaml(path string) {
	yamlData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

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
	fmt.Printf("Detected YAML config File: %s...\n", path)
}

func (config *Config) ReadJson(path string) {
	fmt.Printf("Detected JSON config File: %s...\n", path)

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config Full: %s\n", cfg)
	fmt.Printf("Tap Type: %s\n", cfg.tap.Type)
	fmt.Printf("Tap Name: %s\n", cfg.tap.Name)
	fmt.Printf("Tap Path: %s\n", cfg.tap.Path)
	fmt.Printf("Target Type: %s\n", cfg.target.Type)
	fmt.Printf("Target Name: %s\n", cfg.target.Name)
	fmt.Printf("Detected YAML config File: %s...\n", path)
}
