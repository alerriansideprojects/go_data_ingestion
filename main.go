package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var configPath = "config.yaml"

type Config struct {
	Tap struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"tap"`
	Target struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
	} `yaml:"target"`
}

func main() {
	fmt.Println("Reading config...")
	readConfigFile(configPath)
}

func readConfigFile(cfgPath string) {
	yamlData, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tap Type: %s\n", cfg.Tap.Type)
	fmt.Printf("Tap Name: %s\n", cfg.Tap.Name)
	fmt.Printf("Tap Path: %s\n", cfg.Tap.Path)
	fmt.Printf("Target Type: %s\n", cfg.Target.Type)
	fmt.Printf("Target Name: %s\n", cfg.Target.Name)
}

