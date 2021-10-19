package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var path string
	arguments := os.Args
	if len(arguments) == 1 {
		path = "."
	} else {
		path = os.Args[1]
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {

		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
