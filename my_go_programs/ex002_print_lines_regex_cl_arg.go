package main

// https://gobyexample.com/command-line-arguments
// Run with: go run ex002_print_lines_regex_cl_arg.go "\brs429358\b"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	regex := os.Args[1]
	homeDir := getHomeDir()
	pathToFile := fmt.Sprintf("%s/big_files/1000GENOMES-phase_3.vcf", homeDir)
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(regex)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}
