# Chapter 2 - Basic Go Data Types

## The error data type

Go follows the next convention about error values: if the value of an error variable is nil, then there was no error. 
As an example, let us consider ___strconv.Atoi()___, which is used for converting a string value into an int value (Atoi stands for ASCII to Int).
As specified by its signature, _strconv.Atoi()_ returns (int, error). 
Having an error value of nil means that the conversion was successful and that you can use the int value if you want. 
Having an error value that is not nil means that the conversion was unsuccessful and that the string input is not a valid int value.

"Go follows the next convention about error values" What does this mean

[Good piece on exceptions and errors](https://stackoverflow.com/questions/253314/conventions-for-exceptions-or-error-codes)

"By convention, errors are the last return value and have type error, a built-in interface." See [this link](https://gobyexample.com/errors)

If you want to format your error messages in the way ___fmt.Printf()___ works, you can use the ___fmt.Errorf()___ function, 
which simplifies the creation of custom error messages—the _fmt.Errorf()_ function returns an error value just like ___errors.New()___.

You should have a global error handling tactic in each application that should not change.
 In practice, this means the following:

- All error messages should be handled at the same level, which means that all errors should either be returned to the calling function or be handled at the place they occurred.
- It should be clearly documented how to handle critical errors. This means that there will be situations where a critical error should terminate the program and other times where a critical error might just create a warning message onscreen.
- It is considered a good practice to send all error messages to the log service of your machine because this way the error messages can be examined at a later time. However, this is not always true, so exercise caution when setting this up—for example, cloud native apps do not work that way.

## Numeric data types

Integer data types can be either signed or unsigned, which is not the case for floating point numbers.

```go
package main

import (
	"fmt"
)

func main() {
    c1 := 12 + 1i
    c2 := complex(5, 7)
    fmt.Printf("Type of c1: %T\n", c1)
    fmt.Printf("Type of c2: %T\n", c2)
    var c3 complex64 = complex64(c1 + c2)
    fmt.Println("c3:", c3)
    fmt.Printf("Type of c3: %T\n", c3)
    cZero := c3 - c3
    fmt.Println("cZero:", cZero)
    x := 12
    k := 5
    fmt.Println(x)
    fmt.Printf("Type of x: %T\n", x)
    div := x / k
    fmt.Println("div", div)
    x = 12
    k = 5
    fmt.Println(x)
    fmt.Printf("Type of x: %T\n", x)
    div = x / k
    fmt.Println("div", div)
}
```

## Non-numeric data types

Go has support for Strings, Characters, Runes, Dates, and Times.

### Strings, Characters, and Runes

A Go string is just a collection of bytes and can be accessed as a whole or as an array. 
A single byte can store any ASCII character—however, multiple bytes are usually needed for storing a single Unicode character.

A __rune__ is an int32 value that is used for representing a single Unicode code point, which is an integer value that is used for 
representing single Unicode characters or, less frequently, providing formatting information.

You can create a new byte slice from a given string by using a ___[]byte("A String")___


```go
package main

import (
	"fmt"
)

func main() {
	// You can create a new byte slice from a given string by using a []byte("A String")
	// See the effects of changing %c with %v %s below!
	s := []byte("A String")
	for i, c := range s {
		fmt.Printf("Index %d, element %c\n", i, c)
	}
}
```

Look what happens when you add a unicode character:

```go

package main

import (
	"fmt"
)

func main() {
	// You can create a new byte slice from a given string by using a []byte("A String with a unicode character€")
	s := []byte("A String €")
	for i, c := range s {
		fmt.Printf("Index %d, element %c\n", i, c)
	}
}
```

As strings can be accessed as arrays, you can iterate over the runes of the string using a for loop or point to a specific character if you know
 its place in the string. The length of the string is the same as the number of characters found in the string, which is usually not true for 
byte slices because Unicode characters usually require more than one byte.

```go
package main

import (
	"fmt"
)

func main() {
    aString := "Hello World! €"
    fmt.Println("First character", string(aString[0]))
    // Runes
    // A rune
    r := '€'
    fmt.Println("As an int32 value:", r)
    // Convert Runes to text
    fmt.Printf("As a string: %v and as a character: %c\n", r, r)
    // Print an existing string as runes
    for _, v := range aString {
        fmt.Printf("%x ", v)
    }
    fmt.Println()
    // Print an existing string as characters
    for _, v := range aString {
        fmt.Printf("%c", v)
    }
    fmt.Println()
}
```

### The unicode package

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	sL := "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}

```

### The strings package

The strings standard Go package allows you to manipulate UTF-8 strings in Go 


```go
package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))
	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}

```

### Times and dates

The ___time.Time___ data type, which represents an instant in time with nanosecond precision. Each time.Time value is associated with a location (time zone).

The ___time.Now().Unix()___ function returns the popular UNIX epoch time, which is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970. 
If you want to convert the UNIX time to the equivalent _time.Time_ value, you can use the time.Unix() function. 

#### Age in days function

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	birthday := "2006-01-09"
	today := "2021-10-21"
	daysOld := ageInDays(today, birthday)
	fmt.Printf("Age in days is %0.2f\n", daysOld)
}

func ageInDays(d1ISO, d2ISO string) float64 {
	layoutISO := "2006-01-02"
	d1, err1 := time.Parse(layoutISO, d1ISO)
	d2, err2 := time.Parse(layoutISO, d2ISO)
	if err1 != nil || err2 != nil {
		fmt.Printf("Errors: %v, %v \n", err1, err2)
		os.Exit(1)
	}
	daysOld := d1.Sub(d2).Hours() / 24
	return daysOld
}
```

If you want to parse the 30 January 2020 string and convert it into a Go date variable, you should match it against the 02 January
2006 string—you cannot use anything else in its place when matching a string with the 30 January 2020 format. 
Similarly, if you want to parse the 15 August 2020 10:00 string, you should match it against the 02 January 2006 15:04 string.

you can calculate the time duration between the current time and a time in the past using a call to ___time.Since()___.
Look into this!

#### Working with different time zones

Once again, you need time.Parse() in order to convert a valid input into a time.Time value before doing the conversions. 
This time the input string contains the time zone and is parsed by the "02 January 2006 15:04 MST" string.

In order to convert the parsed date and time into New York time, the program uses the following code:

```go
loc, _ = time.LoadLocation("America/New_York")
fmt.Printf("New York Time: %s\n", now.In(loc))
```

## Go constants

Strictly speaking, the value of a constant variable is defined at compile time, not at runtime—this means that it is included in the binary executable.

___constant generator iota___


```go
const (
        Zero Digit = iota
        One
        Two
        Three
        Four
    )
```

This is the same as:

```go
const (
    Zero = 0
    One = 1
    Two = 2
    Three = 3
    Four = 4
)
```

## Grouping similar data

As a result, arrays in Go are not very powerful, which is the main reason that Go has introduced an additional data structure named slice that is 
similar to an array but is dynamic in nature 

___Slices___ in Go are more powerful than arrays mainly because they are dynamic, which means that they can grow or shrink after creation if needed. 
Additionally, any changes you make to a slice inside a function also affect the original slice. 
But how does this happen? Strictly speaking, all parameters in Go are passed by value—there is no other way to pass parameters in Go.

a slice value is a header that contains a pointer to an underlying array where the elements are actually stored, the length of the array, and 
its capacity—the capacity of a slice is explained in the next subsection. Note that the slice value does not include its elements, 
just a pointer to the underlying array. So, when you pass a slice to a function, Go makes a copy of that header and passes it to the function. 
This copy of the slice header includes the pointer to the underlying array. 

You can create a slice using __make()__ or like an array without specifying its size or using [...]. 
If you do not want to initialize a slice, then using _make()_ is better and faster. 
However, if you want to initialize it at the time of creation, then make() cannot help you. 
As a result, you can create a slice with three float64 elements as aSlice := []float64{1.2, 3.2, -4.5}. Creating a slice with space 
for three float64 elements with make() is as simple as executing make([]float64, 3). Each element of that slice has a value of 0, 
which is the zero value of the float64 data type.

You can find the length of an array or a slice using ___len()___. 
As you will find out in the next subsection, slices have an additional property named capacity. 
You can add new elements to a full slice using the append() function. append() automatically allocates the required memory space.

```go
package main
import "fmt"
func main() {
    // Create an empty slice
    aSlice := []float64{}
    // Both length and capacity are 0 because aSlice is empty
    fmt.Println(aSlice, len(aSlice), cap(aSlice))
    // Add elements to a slice
    aSlice = append(aSlice, 1234.56)
    aSlice = append(aSlice, -34.0)
    fmt.Println(aSlice, "with length", len(aSlice))
    // A slice with length 4
    t := make([]int, 4)
    t[0] = -1
    t[1] = -2
    t[2] = -3
    t[3] = -4
    // Now you will need to use append
    t = append(t, -5)
    fmt.Println(t)
    // A 2D slice
    // You can have as many dimensions as needed
    twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
    // Visiting all elements of a 2D slice
    // with a double for loop
    for _, i := range twoD {
            for _, k := range i {
                fmt.Print(k, " ")
            }
            fmt.Println()
    }
    make2D := make([][]int, 2)
    fmt.Println(make2D)
    make2D[0] = []int{1, 2, 3, 4}
    make2D[1] = []int{-1, -2, -3, -4}
    fmt.Println(make2D)
}
```

The capacity shows how much a slice can be expanded without the need to allocate more memory and change the underlying array. 
Although after slice creation the capacity of a slice is handled by Go, a developer can define the capacity of a slice at creation 
time using the make() function—after that the capacity of the slice doubles each time the length of the slice is about to become bigger 
than its current capacity. 


But what happens when you want to append a slice or an array to an existing slice? Should you do that element by element? 
Go supports the ___... operator___, which is used for exploding a slice or an array into multiple arguments before appending it to an existing slice.

Setting the correct capacity of a slice, if known in advance, will make your programs faster because Go will not have to allocate a new underlying 
array and have all the data copied over.


A __byte slice__ is a slice of the byte data type ([]byte). 
Go knows that most byte slices are used to store strings and so makes it easy to switch between this type and the string type. 
There is nothing special in the way you can access a byte slice compared to the other types of slices.
What is special is that Go uses byte slices for performing file I/O operations because they allow you to determine with precision the 
amount of data you want to read or write to a file. 
This happens because bytes are a universal unit among computer systems.














