# Chater 5 - Go Packages and Functions

Packages are Go's way of organizing, delivering, and using code.

Go also supports modules, which are packages with version numbers.

__defer__  is used for cleaning up and releasing resources.

Go follows a simple rule that states that functions, variables, data types, structure fields, and so forth that begin with an uppercase letter are __public__,
 whereas functions, variables, types, and so on that begin with a lowercase letter are __private__. 

The same rule applies not only to the name of a struct variable but to the fields of a struct variable—in practice, this means that you 
 can have a struct variable with both private and public fields

## Go packages

Everything in Go is delivered in the form of packages.

A Go package is a Go source file that begins with the __package keyword__, followed by the name of the package

Note that packages can have structure. For example, the __net__ package has several subdirectories,
 named _http, mail, rpc, smtp, textproto_, and _url_, which should be imported as:
 _net/http, net/mail, net/rpc, net/smtp, net/textproto_, and _net/url_, respectively.

Packages are mainly used for grouping related functions, variables, and constants so that you can transfer them easily and use them in your own Go programs. 
Note that apart from the _main_ package, Go packages are not autonomous programs and cannot be compiled into executable files on their own.

Tried installing cobra as per book:

```{console}
go get github.com/spf13/cobra
# Install path on my machine:
~/go/pkg/mod/cache/download/github.com/spf13
```

Starting with Go 1.16, go install is the recommended way of building and installing packages in module mode. The use of go get is deprecated

## Diversion

Worked through [this tutorial](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177)

To upgrade an existing package, you should execute go get with the __-u__ option. 
Add the __-v__ option to the _go get_ command to see what is happenning behind the scenes

## Functions

The single-most popular Go function is __main()__, which is used in every executable Go program—the main() function accepts no parameters and returns nothing,
 but it is the starting point of every Go program. 
When the main() function ends, the entire program ends as well.

__Anonymous functions__ can be defined inline without the need for a name, and they are usually used for implementing things that require a small amount of code. 

A function can return an anonymous function or take an anonymous function as one of its arguments. 
Anonymous functions can be attached to Go variables.
Called __lambdas__ in functional programming terminology.
A __closure__ is a specific type of anonymous function that carries or closes over variables that are in the same lexical scope as the anonymous function that was defined.


Functions can return multiple values

```go
package main

import (
	"fmt"
)

func main() {
	a, b, c, d := multi(10)
	fmt.Printf("%d %d %d %d\n", a, b, c, d)
}

func multi(x int) (int, int, int, int) {
	return x * x, x + x, x - x, x/x
}
```

Note the compulsory use of parentheses when a function returns more than one value.

The return values of a function can be __named__

Functions can accept other functions as parameters. 
The best example of a function that accepts another function as an argument can be found in the __sort__ package

You are not obliged to use an anonymous function in either _sort.Slice()_ or _sort.SliceIsSorted()_. 
You can define a regular function with the required signature and use that.

```go
package main
import (
    "fmt"
    "sort"
)
type Grades struct {
    Name    string
    Surname string
    Grade   int
}
func main() {
    data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
        {"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}
    isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
        return data[i].Grade < data[j].Grade
    })
    if isSorted {
        fmt.Println("It is sorted!")
    } else {
        fmt.Println("It is NOT sorted!")
    }
    sort.Slice(data,
        func(i, j int) bool { return data[i].Grade < data[j].Grade })
    fmt.Println("By Grade:", data)
}
```

__Functions can return other functions__

Apart from accepting functions as arguments, functions can also return anonymous functions, which can be handy when the returned function is not always
 the same but depends on the function's input or other external parameters. 


