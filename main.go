package main

import (
	"fmt"
	"os"
	config "go_data_ingestion/bin"
)

func main() {
	fmt.Println("Reading config...")
	if files, err := os.ReadDir("./configs/"); err != nil {
		panic(err)
	} else {
		for _, file := range files {
			configPath := "./configs/" + file.Name()
			var c = config.Config{}
			switch configPath[len(configPath)-4:] {
			case "yaml":
				c.ReadYaml(configPath)
			case "json":
				c.ReadJson(configPath)
			}
		}
	
	}

}
