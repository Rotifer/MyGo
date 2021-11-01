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
	columnValues := uniqueListFromFileColumn("./data/analysis.csv", 8)
	elementsInMap := getElementsInMap(columnValues, dict)
	fmt.Printf("%v\n", columnValues)
	fmt.Println("%v\n", elementsInMap)
}

func buildPubmedIDToStudyAccessionMap(filePath string) map[string]string {
	pubmedIdToStudyaccessionMap := make(map[string]string)
	_, err := os.Stat(filePath)
	if err != nil {
		os.Exit(1)
	}
	f, err := os.Open(filePath)
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

func  uniqueListFromFileColumn(filePath string, columnIndex int) []string {
	var columnValues  []string
	var value string
	setMap := make(map[string]string)
	_, err := os.Stat(filePath)
	if err != nil {
		os.Exit(1)
	}
	
	f, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		value = record[columnIndex]
		setMap[value] = value
	}

	for k, _ := range setMap {
		columnValues = append(columnValues, k)
	}

	return columnValues
}


func getElementsInMap(s []string, dict map[string]string) []string {
	var elementsInMap []string
	for _, el := range s {
		if _, ok :=  dict[el]; ok {
			elementsInMap = append(elementsInMap, el)
		}
	}
	return elementsInMap
}
