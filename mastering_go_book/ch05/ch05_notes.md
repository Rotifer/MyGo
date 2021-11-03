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



