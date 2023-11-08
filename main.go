package main

import (
	"fmt"
	config "go_data_ingestion/bin"
)

var configPath = "./configs/config.yaml"

func main() {
	fmt.Println("Reading config...")
	var c = config.Config{}
	c.Read(configPath)
}

