package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var err = filepath.Walk("./", func(file string, _ os.FileInfo, _ error) error {
		fmt.Println("__________")
		basename := filepath.Base(file)
		fmt.Println(basename)
		fp := filepath.Dir(file)
		fmt.Println(fp)

		return nil
	})
	if err != nil {
		fmt.Println("Error: %+v\n", err)
	}
}
