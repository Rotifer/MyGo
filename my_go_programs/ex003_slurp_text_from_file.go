package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "ex003_lorem_ipsum.txt"
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	text, err := io.ReadAll(file)
	fmt.Println(string(text))
}



