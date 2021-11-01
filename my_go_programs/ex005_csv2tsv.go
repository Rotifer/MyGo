package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func main() {
	inputCsvFilePath := "./data/analysis.csv"
	csv2tsv(inputCsvFilePath)
	fmt.Println("Done!")
}

func csv2tsv(csvFilePath string) {
	_, err := os.Stat(csvFilePath)
	if err != nil {
		os.Exit(1)
	}

	f, err := os.Open(csvFilePath)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	tsvFilePath := strings.TrimSuffix(csvFilePath, path.Ext(csvFilePath)) + ".tsv"
	fo, err := os.Create(tsvFilePath)
	defer fo.Close()
	writer := csv.NewWriter(fo)
	writer.Comma = '\t'
	defer writer.Flush()
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		writer.Write(record)
	}

}
