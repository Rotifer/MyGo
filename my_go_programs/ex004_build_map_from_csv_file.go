package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/Users/mfm45656/gwas_catalog/gwas-catalog-download-ancestries-v1.0.3.txt"
	dict := buildPubmedIDToStudyAccessionMap(filePath)
	fmt.Println(len(dict))
	key := "24954895"
	fmt.Println(dict[key])
}

func buildPubmedIDToStudyAccessionMap(filepath string) map[string]string {
	pubmedIdToStudyaccessionMap := make(map[string]string)
	_, err := os.Stat(filepath)
	if err != nil {
		os.Exit(1)
	}
	f, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.Comma = '\t'
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		pubmedIdToStudyaccessionMap[record[1]] = record[0]
	}
	return pubmedIdToStudyaccessionMap
}

