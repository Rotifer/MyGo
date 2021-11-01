# Chapter 4 Reflection and Interfaces

Interfaces are about expressing abstractions and identifying and defining behaviors that can be shared among different data types. 

Interfaces work with __methods on types__ or type methods, which are like functions attached to given data types, which in Go are usually structures. 

## Reflection

How you can find out the names of the fields of a structure at execution time? 
You need to use reflection. 
Apart from enabling you to print the fields and the values of a structure, reflection also allows you to explore and manipulate unknown structures
 like the ones created from decoding JSON data

Go provides the __reflect__ package for working with reflection.

The introduction of __generics__ in Go might make the use of reflection less frequent in some cases, because with generics you can work with
 different data types more easily and without the need to know their exact data types in advance.

The most useful parts of the reflect package are __two data types__ named __reflect.Value__ and __reflect.Type__. 
Now, _reflect.Value_ is used for storing values of any type, whereas _reflect.Type_ is used for representing Go types. 

There are two functions:

1. reflect.TypeOf() returns _reflect.Type_
2. reflect.ValueOf() returns _reflect.Value_

As structures are really important in Go, the reflect package offers the __reflect.NumField()__ method for listing the number
 of fields in a structure as well as the __Field()__ method for getting the __reflect.Value__ value of a specific field of a structure.

```go
package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}
type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())
		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
```

Simple _reflect_ example

```go
package main

import (
	"fmt"
	"reflect"
)



func main() {
	s := []string{"one", "two"}
	r := reflect.ValueOf(s)
	t := r.Type()
	fmt.Printf("%s\n", t.String())
}
```

## Three disadvantages of reflection

1. Readability/maintenance
2. Speed
3. Errors not caught in build time but at runtime so cause runtime panics

## Type methods


