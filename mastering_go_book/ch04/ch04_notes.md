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

A __type method__ is a function that is attached to a specific data type.

Defining new type methods is as simple as creating new functions, provided that you follow certain rules that
 associate the function with a data type.

```go
func (receiver object) FunctionName(parameters) <return values> {
    ...
}
```

It is the __receiver object__ variable in parentheses between _func_ and the function name that turns the function into a __type method__

Te Go compiler does turn methods into regular function calls with the __self__ value as the first parameter. 
Interfaces require the use of type methods to work.

## Interfaces

An __interface__ is a Go mechanism for defining behavior that is implemented using a set of methods. 
Interfaces play a key role in Go and can simplify the code of your programs when they have to deal with multiple data
 types that perform the same task

If you decide to create your own interfaces, then you should begin with a common behavior that you want to be used by multiple data types.

Once you implement the required type methods of an interface, that interface is satisfied __implicitly__.

The __empty interface__ is defined as just _interface{}_. 
As the empty interface has no methods, it means that it is already implemented by all data types.

A Go interface type defines (or describes) the __behavior__ of other types by specifying a set of methods that need to be implemented
 for supporting that behavior. 
For a data type to satisfy an interface, it needs to implement all the type methods required by that interface. 
Therefore, interfaces are __abstract types__ that specify a set of methods that need to be implemented so that another type can be
 considered an instance of the interface.So, an interface is two things: 
1. a set of methods
2. a type. 


The biggest advantage you get from interfaces is that if needed, you can pass a variable of a data type that implements a particular interface to any function
 that expects a parameter of that specific interface, which saves you from having to write separate functions for each supported data type. 
However, Go offers an alternative to this with the __recent addition of generics__.

Interfaces can also be used for providing a kind of __polymorphism__ in Go, which is an object-oriented concept. 
Polymorphism offers a way of accessing objects of different types in the same uniform way when they share a common behavior.

Interfaces can be used for __composition__. 
In practice, this means that you can combine existing interfaces and create new ones that offer the combined behavior of the interfaces that were brought together.

### The _sort.Interface_ interface

The sort package contains an interface named sort.Interface that allows you to sort slices according to your needs and your data, provided
 that you implement _sort.Interface_ for the custom data types stored in your slices. 

In order to __implement sort.Interface__, we need to implement the following three type methods:

1. Len() int
2. Less(i, j int) bool
3. Swap(i, j int)


```go
package main
import (
    "fmt"
    "sort"
)
type S1 struct {
    F1 int
    F2 string
    F3 int
}
// We want to sort S2 records based on the value of F3.F1,
// Which is equivalent to S1.F1 as F3 is an S1 structure
type S2 struct {
    F1 int
    F2 string
    F3 S1
}
type S2slice []S2

// Implementing sort.Interface for S2slice
func (a S2slice) Len() int {
    return len(a)
}

// What field to use when comparing
func (a S2slice) Less(i, j int) bool {
    return a[i].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func main() {
    data := []S2{
        S2{1, "One", S1{1, "S1_1", 10}},
        S2{2, "Two", S1{2, "S1_1", 20}},
        S2{-1, "Two", S1{-1, "S1_1", -20}},
    }
    fmt.Println("Before:", data)
    sort.Sort(S2slice(data))
    fmt.Println("After:", data)
    // Reverse sorting works automatically
    sort.Sort(sort.Reverse(S2slice(data)))
    fmt.Println("Reverse:", data)
}
```

### The empty interface

A function with an _interface{}_ parameter can accept variables of any data type in this place. 
However, if you intend to work with _interface{}_ function parameters without examining their data type inside the function,
 you should process them with statements that work on all data types, otherwise your code may crash or misbehave.

### Type assertions and type switches

A type assertion is a mechanism for working with the underlying concrete value of an interface. 

__Type switches__ use __switch blocks__ for data types and allow you to differentiate between type assertion values, which are data types,
 and process each data type the way you want. 
Im order to use the empty interface in type switches, you need to use type assertions.

Type assertions use the __x.(T)__ notation, where x is an interface type and T is a type, and help you extract the value that is hidden behind the empty interface. 
For a type assertion to work, x should not be nil and the dynamic type of x should be identical to the T type.

```go
package main
import "fmt"

type Secret struct {
    SecretValue string
}
type Entry struct {
    F1 int
    F2 string
    F3 Secret
}
func Teststruct(x interface{}) {
    // type switch
    switch T := x.(type) {
    case Secret:
        fmt.Println("Secret type")
    case Entry:
        fmt.Println("Entry type")
    default:
        fmt.Printf("Not supported type: %T\n", T)
    }
}
func Learn(x interface{}) {
    switch T := x.(type) {
    default:
        fmt.Printf("Data type: %T\n", T)
    }
}
func main() {
    A := Entry{100, "F2", Secret{"myPassword"}}
    Teststruct(A)
    Teststruct(A.F3)
    Teststruct("A string")
    Learn(12.23)
    Learn('€')
}
```

### The map[string]interface{} map

The biggest advantage you get from using a _map[string]interface{_} map or any map that stores an _interface{_} value in general,
 is that you still have your data in its original state and data type.

The _map[string]interface{}_ is good at storing JSON records of an unknown type.

```go
package main
import (
    "encoding/json"
    "fmt"
    "os"
)
var JSONrecord = `{
    "Flag": true,
    "Array": ["a","b","c"],
    "Entity": {
      "a1": "b1",
      "a2": "b2",
      "Value": -456,
      "Null": null
    },
    "Message": "Hello Go!"
  }`

func typeSwitch(m map[string]interface{}) {
    for k, v := range m {
        switch c := v.(type) {
        case string:
            fmt.Println("Is a string!", k, c)
        case float64:
            fmt.Println("Is a float64!", k, c)
        case bool:
            fmt.Println("Is a Boolean!", k, c)
        case map[string]interface{}:
            fmt.Println("Is a map!", k, c)
            typeSwitch(v.(map[string]interface{}))
        default:
            fmt.Printf("...Is %v: %T!\n", k, c)
        }
    }
    return
}

// Recursion here
func exploreMap(m map[string]interface{}) {
    for k, v := range m {
        embMap, ok := v.(map[string]interface{})
        // If it is a map, explore deeper
        if ok {
            fmt.Printf("{\"%v\": \n", k)
            exploreMap(embMap)
            fmt.Printf("}\n")
        } else {
            fmt.Printf("%v: %v\n", k, v)
        }
    }
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println("*** Using default JSON record.")
    } else {
        JSONrecord = os.Args[1]
    }
    JSONMap := make(map[string]interface{})
    err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
    if err != nil {
        fmt.Println(err)
        return
    }
    exploreMap(JSONMap)
    typeSwitch(JSONMap)
}

```

### The error data type

The error data type is an interface defined as follows:

```go
type error interface {
    Error() string
}
```

In order to satisfy the error interface you just need to implement the __Error()__ string type method.

When there is nothing more to read from a file, Go returns an __io.EOF error__, which, strictly speaking, 
 is not an error condition but a logical part of reading a file. 
If a file is totally empty, you still get __io.EOF__ when you try to read it. 

### Writing your own interfaces

We usually want to share our interfaces, which means that interfaces are usually included in Go packages other than _main_.

The fact that Go considers interfaces as data types allows us to create slices with elements that satisfy a given interface without getting any error messages.

Once you have implemented sort.Interface, you can also sort your data in reverse order using __sort.Reverse()__.

## Working with two different CSV file formats

## Object-oriented programming in Go

As Go does not support all object-oriented features, it cannot replace an object-oriented programming language fully. 
However, it can mimic some object-oriented concepts.

1. A Go structure with its type methods is like an object with its methods. 
2. Interfaces are like abstract data types that define behaviors and objects of the same class, which is similar to polymorphism. 
3. Go supports encapsulation, which means it supports hiding data and functions from the user by making them private to the structure and the current Go package. 
4. Combining interfaces and structures is like composition in object-oriented terminology.


Note on Environment variables

```go
filepath := os.Getenv("PHONEBOOK")
```

## Exercises

See example: MyGo/my_go_programs/ex006_sort_slice.go
