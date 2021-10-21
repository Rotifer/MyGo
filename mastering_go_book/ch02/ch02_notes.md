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

















