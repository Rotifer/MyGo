# Chapter 6 - Telling a UNIX System What to Do

__Important note:__ 
Starting with Go 1.16, the GO111MODULE environment variable defaults to on—this affects the use of Go packages that do not belong to the Go standard library. 
In practice, this means that you must put your code under _~/go/src_. 
You can go back to the previous behavior by setting GO111MODULE to auto, but you do not want to do that—modules are the future. 
The reason for mentioning this in this chapter is that __both viper and cobra__ prefer to be treated as Go modules instead of packages,
 which changes the development process but not the code.

## stdin, stdout, and stderr

By default, all UNIX systems support three special and standard filenames: 
1. /dev/stdin
2. /dev/stdout
3. /dev/stderr
which can also be accessed using file descriptors 0, 1, and 2, respectively.

Go uses:
__os.Stdin__ for accessing standard input, 
__os.Stdout__ for accessing standard output
__os.Stderr__ for accessing standard error. 
Although you can still use /dev/stdin, /dev/stdout, and /dev/stderr or the related file descriptor values for accessing the same devices,
 it is better, safer, and more portable to stick with _os.Stdin, os.Stdout_, and _os.Stderr_.

## UNIX processes

There are three process categories: 
1. user processes
2. daemon processes
3. kernel processes

Although you can fork a new process in Go using the __exec__ package, Go does not allow you to control threads
Go offers __goroutines__, which the user can create on top of threads that are created and handled by the Go runtime.

## Handling UNIX signals

UNIX signals offer a very handy way of interacting asynchronously with your applications. (what way?)

A __goroutine__ is the smallest executable Go entity. 
In order to create a new goroutine you have to use the _go_ keyword followed by a predefined function or an anonymous function—the methods are equivalent. 
A __channel__ in Go is a mechanism that among other things allows goroutines to communicate and exchange data.

Read the section on signals but made no notes. Will return to this

## File I/O

