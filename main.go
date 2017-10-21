package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var err = filepath.Walk("./", func(file string, _ os.FileInfo, _ error) error {
		basename := filepath.Base(file)
		fmt.Println("Basename: ", basename)
		fp := filepath.Dir(file)
		fmt.Println("Dir: ", fp)
		fmt.Println("__________")

		return nil
	})
	if err != nil {
		fmt.Println("Error: %+v\n", err)
	}
}
