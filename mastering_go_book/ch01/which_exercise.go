package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	path := os.Getenv("PATH")
	fmt.Printf("PATH: %s\n", path)
	pathSplit := filepath.SplitList(path)
	for _, file := range arguments {
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						fmt.Println(fullPath)
						//return
					}
				}
			}
		}
	}
}
