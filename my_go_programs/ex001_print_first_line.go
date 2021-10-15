package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	homeDir := getHomeDir()
	pathToFile := fmt.Sprintf("%s/big_files/1000GENOMES-phase_3.vcf", homeDir)
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstLine string
	for scanner.Scan() {
		firstLine = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(firstLine)
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}
