package main

/*
Links:
https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field/42872183#42872183
https://www.dotnetperls.com/struct-go
https://zetcode.com/golang/struct/
*/
import (
	"fmt"
	"sort"
)

type SNP struct {
	ID         string
	chromosome string
	position   int
	refAllele  string
	altAllele  string
}

func main() {
	var snps []SNP
	snp1 := new(SNP)
	snp1.ID = "rs429358"
	snp1.chromosome = "19"
	snp1.position = 44908684
	snp1.refAllele = "T"
	snp1.altAllele = "C"

	snp2 := new(SNP)
	snp2.ID = "rs7412"
	snp2.chromosome = "19"
	snp2.position = 44908822
	snp2.refAllele = "C"
	snp2.altAllele = "T"
	snp3 := new(SNP)
	snp3.ID = "rs405509"
	snp3.chromosome = "19"
	snp3.position = 44905579
	snp3.refAllele = "A"
	snp3.altAllele = "C"
	snps = append(snps, *snp1)
	snps = append(snps, *snp2)
	snps = append(snps, *snp3)
	fmt.Println("Before sort:")
	fmt.Printf("%v\n", snps)
	sort.Slice(snps, func(i, j int) bool {
		return snps[i].position < snps[j].position
	})
	fmt.Println("After sort:")
	fmt.Printf("%v\n", snps)
	sort.Slice(snps, func(i, j int) bool {
		return snps[i].position > snps[j].position
	})	
	fmt.Println("After reverse sort:")
	fmt.Printf("%v\n", snps)
}

