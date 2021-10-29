# Composite Data Types

- Maps
- Structures
- Pointers and structures
- Regular expressions and pattern matching

## Maps

Maps allow you to use indexes of various data types as keys to look up your data as long as these keys are __comparable__.

Although floating point values are comparable, precision issues caused by the internal representation of such values might create bugs and crashes, 
so you might want to avoid using floating point values as keys to Go maps.

- Maps are very versatile. 
- Although this is not always the case, working with maps in Go is fast, as you can access all elements of a map in __linear__ time. 
- Inserting and retrieving elements from a map is fast and does not depend on the __cardinality__ of the map.
- Maps are easy to understand, which leads to clear designs.

```go
package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{}
	m3 := make(map[rune]int)
	m1["a"] = 97
	m2["A"] = 65
	m3['b'] = 98
	fmt.Printf("Map1: %v ; Map2: %v ; Map3: %v\n", m1, m2, m3)
	k, ok := m3['c']
	fmt.Println(k, ok)
}
```

Output from above: Map1: map[a:97] ; Map2: map[A:65] ; Map3: map[98:98]

- You can find the length of a map, which is the number of keys in the map, using the __len()__ function.
- You can delete a key and value pair from a map using the __delete()__ function, which accepts two arguments: 
  the name of the map and the name of the key, in that order
- Does a key exist in a map? _v, ok := aMap[k]_ No error is raised if the key does not exists. If it does exist, _ok_ is set to _true_

### Iterating over maps

When __for__ is combined with the __range__ keyword it implements the functionality of _foreach_ loops found in other programming languages 
 and allows you to iterate over all the elements of a map without knowing its size or its keys. 
When _range_ is applied on a map, it returns key and value pairs in that order.


## Structures

Structures, as well as other user-defined data types, are usually defined outside the _main()_ function or any other package function so that they have a
 global scope and are available to the entire Go package. 
Therefore, unless you want to make clear that a type is only useful within the current local scope and is not expected to be used elsewhere,
 you should write the definitions of new data types outside functions.

When you define a new structure, you group a set of values into a single data type, which allows you to pass and receive this set of values as a single entity. 

A structure has __fields__, and each field has its own data type, which can even be another structure or slice of structures. 
Additionally, as a structure is a new data type, it is defined using the __type__ keyword followed by the name of the structure
 and ending with the __struct__ keyword, which signifies that we are defining a new structure.

There are two ways to work with structure variables:
1. As regular variables
2. As pointer variables that point to the memory address of a structure.

__NB__:
The __order__ in which you put the fields in the definition of a structure type __is significant__ for the type identity of the defined structure. 
Two structures with the same fields will not be considered identical in Go if their fields are not in the same order.

Can create new structure instances using the __new()__ keyword: _pS := new(Entry)_. 
The _new()_ keyword has the following properties:

- It allocates the proper memory space, which depends on the data type, and then it zeroes it
- It always __returns a pointer__ to the allocated memory
- It works for all data types except _channel_ and _map_

_NB_:
An important Go rule: 
If no initial value is given to a variable, the Go compiler automatically __initializes__ that variable to the __zero value of its data type__. 

Generally speaking, when you have __to initialize lots of structure variables__, 
 it is considered a good practice to __create a function__ for doing so as this is less error-prone.

## Regular expressions and pattern matching

The Go package responsible for defining regular expressions and performing pattern matching is called __regexp__. 
We use the __regexp.MustCompile()__ function to create the regular expression and the __Match()__ function to see whether the given string is a match or not.

The _re.Match()_ method returns _true_ if the given byte slice matches the regular expression, which is a _regexp.Regexp_ variable, and _false_ otherwise.

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	testDates := []string{"2021-10-20", "1964-05-21", "973-01-09"}
	for _, d := range testDates {
		fmt.Println(IsYYYYMMDDFormat(d))
	}
}

func IsYYYYMMDDFormat(s string) bool {
	sAsByteSlice := []byte(s)
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.Match(sAsByteSlice)
}
```

Notes on the code:

- The string to match against is turned into a __byte slice__
- The regex argument to __MustCompile__ is enclosed in back ticks

## Reading a tab-separated file and building a data structure from it

- Taking a GWAS Catalog file as input
- File is tab-separated
- The columns of interst are the first two: STUDY ACCESSION	PUBMED ID
- These values repeat - not sure if the relationship is 1:1, that is each study accession is associated with just one PubMed ID and that the inverse is also true
- Objective: Read through the file and build a map of slices where the keys are the PubMed IDs and the values are the study accessions


