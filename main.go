package main

import (
	"fmt"
	"log"
	"os"
	helpers "helpers"
)

func main() {
	files, err := os.ReadDir("./csv_files")
	if err != nil {
		log.Fatal(err)
	}

	helpers.ReadFiles(files)
	// for _, file := range(files) {
	// 	fmt.Println(file.Name())
	// }
}
