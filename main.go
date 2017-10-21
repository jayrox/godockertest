package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var err = filepath.Walk("./", func(file string, _ os.FileInfo, _ error) error {
		fmt.Println(file)
		return nil
	})
	if err != nil {
		fmt.Println("Error: %+v\n", err)
	}
}
