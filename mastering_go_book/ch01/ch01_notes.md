# Chapter 1



There are two ways to execute Go code: 

1. as a compiled language using go build
2. or as a scripting language using go run. 

The _go run_ command builds the named Go package, which in this case is the main package implemented in a single file, 
creates a temporary executable file, executes that file, and deletes it once it is done—to our eyes, this looks like using a scripting language. 


## Variables

__Very important Go rule__: 
if no initial value is given to a variable, the Go compiler will automatically initialize that variable to the zero value of its data type.

There is also the := notation, which can be used instead of a var declaration. := defines a new variable by inferring the data of the value that follows it. 
The official name for := is __short assignment statement__ and it is very frequently used in Go, especially for getting the return values from 
functions and for loops with the range keyword.

The __var__ keyword is mostly used for declaring global or local variables without an initial value. 

The short assignment statement cannot be used outside of a function environment because it is not available there. 

only _const_  and _var_ work for global variables

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 22/7.0
	fmt.Printf("PI: %0.2f\n", pi)
	fmt.Printf("Absolute PI: %0.2f\n", math.Abs(pi * -1))
}
```

### Code Notes

- Show how to print formatted values
- Use of _math_ package
- Note: the expression 22/7 returns an _integer_ type; to return a _float_, you need to set one of the operands as a _float_ type



Go does not allow implicit data conversions like C.

For conversions that are not straightforward (for example, string to int), there exist specialized functions that allow you to catch issues 
with the conversion in the form of an error variable that is returned by the function.

## Controlling program flow

A very common pattern in Go that is used almost everywhere. This pattern says that if the value of an error variable as returned from a function is nil, 
then everything is OK with the function execution. Otherwise, there is an error condition somewhere that needs special care.

The ___switch___ statement has two different forms. In the first form, the switch statement has an expression that is being evaluated, whereas in the second form, 
the _switch_ statement has no expression to evaluate. In that case, expressions are evaluated in each case statement, which increases the flexibility of switch. 

## Iterating with for loops and range

Go supports ___for___ loops as well as the ___range___ keyword for iterating over all the elements of arrays, slices, and  maps.

### The traditional C-style _for_loop

- Initiation, Condition, Increment

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// countdown
	fmt.Println("Commencing countdown.......")
	for i := 10; i > 0; i-- {
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Blast-off!")
}
```

### The _foreach_ style of _for_ loop for iterating over a _slice_

```go
package main

import (
	"fmt"
)

func main() {
	animals := []string{"cat", "dog", "rat", "cow"} // Declare a slice
	for i, animal := range animals {
		fmt.Printf("Index: %d => Animal %s\n", i, animal)
	}
}
```

## Getting user input

The ___fmt.Scanln()___ function can help you read user input while the program is already running and store it to a string variable, which is passed as a pointer 
to _fmt.Scanln()_. The fmt package contains additional functions for reading user input from the console (_os.Stdin_), from files or from argument lists.

## Working with command-line arguments

Usually, user input is given in the form of command-line arguments to the executable file. 

By default, command-line arguments in Go are stored in the ___os.Args___ slice. 
Go also offers the ___flag___ package for parsing command-line arguments, but there are better and more powerful alternatives.

The _os.Args_ slice is properly initialized by Go and is available to the program when referenced. 
The _os.Args_ slice contains string values

The __first__ command-line argument stored in the _os.Args_ slice is always __the name of the executable__. 
If you are using go run, you will get a temporary name and path, otherwise, it will be the path of the executable as given by the user. 
The remaining command-line arguments are what comes after the name of the executable—the various command-line arguments are automatically separated by space 
characters unless they are included in double or single quotes.

## Understanding the Go concurrency model

In order to create a new __goroutine__, you have to use the go keyword followed by a predefined function or an anonymous function—both methods are equivalent 
as far as Go is concerned.

A __channel__ in Go is a mechanism that, among other things, allows goroutines to communicate and exchange data. 

Although goroutines are do not share any variables, they can share memory.

the ___go___ keyword is used for creating goroutines.

```go
package main
import (
    "fmt"
    "time"
)
func myPrint(start, finish int) {
    for i := start; i <= finish; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()
    time.Sleep(100 * time.Microsecond)
}
func main() {
    for i := 0; i < 5; i++ {
        go myPrint(i, 5)
    }
    time.Sleep(time.Second)
}
```

If you run it multiple times, you'll most likely get a different output each time.
This happens because goroutines are initialized in random order and start running in random order. 
The Go scheduler is responsible for the execution of goroutines just like the OS scheduler is responsible for the execution of the OS threads.

## Developing the which(1) utility in Go

See _which.go_

First, we read the command-line arguments of the program (os.Args) and save the first command-line argument into the file variable. 
Then, we get the contents of the ___PATH___ environment variable and split it using ___filepath.SplitList()___, which offers a portable 
way of separating a list of paths. 
Lastly, we iterate over all the directories of the PATH variable using a for loop with range as filepath.SplitList() returns a slice.

We construct the full path that we examine using filepath.Join() that is used for concatenating the different parts of a path using
 an OS-specific separator—this makes filepath.Join() work in all supported operating systems. In this part, we also get some lower-level 
information about the file—remember that in UNIX everything is a file, which means that we want to make sure that we are dealing with a 
regular file that is also executable.

## Logging information

Generally speaking, using a log file to write some information used to be considered a better practice than writing the same output on screen for two reasons: 
1. firstly, because the output does not get lost as it is stored on a file, 
2. and secondly, because you can search and process log files using UNIX tools, such as grep(1), awk(1), and sed(1), which cannot be done when messages are printed on a terminal window. 

However, this is not true anymore.

As we usually run our services via systemd, programs should log to stdout so systemd can put logging data in the journal. 
https://12factor.net/logs offers more information about app logs. Additionally, in cloud native applications, 
we are encouraged to simply log to stderr and let the container system redirect the stderr stream to the desired destination.

The ___log___ package sends log messages to standard error. 
Part of the log package is the _log/syslog_ package, which allows you to send log messages to the syslog server of your machine. 
Although by default log writes to standard error, the use of _log.SetOutput(_) modifies that behavior. 
The list of functions for sending logging data includes _log.Printf()_, _log.Print()_, _log.Println()_, _log.Fatalf()_, _log.Fatalln()_, log.Panic(), log.Panicln() and log.Panicf().

In order to write to system logs, you need to call the ___syslog.New()___ function with the appropriate parameters. Writing to the main system log file 
is as easy as calling syslog.New() with the syslog.LOG_SYSLOG option. 
After that you need to tell your Go program that all logging information goes to the new logger—this is implemented with a call to the ___log.SetOutput()___ function.


### log.Fatal() and log.Panic()

The ___log.Fatal()___ function is used when something erroneous has happened and you just want to exit your program as soon as possible after reporting that bad situation
The call to log.Fatal() terminates a Go program at the point where log.Fatal() was called after printing an error message.
In most cases, this custom error message can be Not enough arguments, Cannot access file, or similar. 
Additionally, it returns back a non-zero exit code, which in UNIX indicates an error. 

log.Panic() is equivalent to a call to log.Print() followed by a call to panic(). panic() is a built-in function that stops the execution of the current function and begins panicking. 

### Writing to a custom log file

Most of the time, and especially on applications and services that are deployed to production, you just need to write your logging data in a log file of your choice. 

See program __custom_log.go__

The ___defer___ keyword tells Go to execute the statement just before the current function returns. 
This means that f.Close() is going to be executed just before main() returns.

The last three statements create a new log file based on an opened file (f) and write two messages to it using Println().

Getting the Temp Dir path:


```go
package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	tmpDir := path.Join(os.TempDir())
	fmt.Printf("Temp Dir: %v\n", os.TempDir())
	fmt.Printf("The temp dir on this machine is %s\n", tmpDir)
}
```

## Overview of Go generics

Slated for version 1.18

The main idea behind generics in Go, as well as any other programming language that supports generics, is not having to write special code for 
supporting multiple data types when performing the same task.

However, demanding every programmer to write lots of code and implement lots of functions and methods for supporting multiple custom data types 
is not the optimal solution—generics comes into play for providing an alternative to the use of interfaces and reflection for supporting multiple data types. 

## Developing a basic phone book application


