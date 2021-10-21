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
	s := []byte("A String")
	for i, c := range s {
		fmt.Printf("Index %d, element %c\n", i, c)
	}
}
```



