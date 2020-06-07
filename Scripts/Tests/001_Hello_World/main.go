/*
Lokum GoLang
Learn Go Programming Language with small programs examples
----------
by Walid Amriou
https://lokumgolang.walidamriou.com
----------
How to use:
- To run the program use go run:
$ go run helloworld.go
hello to world of Go Language!
You know that 1+1 is: 2

- To build it to binaries. use go build:
$ go build helloworld.go
$ ls
helloworld  helloworld.go

- To execute the built binary after build it use ./helloworld:
$ ./helloworld
hello to world of Go Language!
----------
*/

/*
Go code is organize d into packages, which are similar to libraries or modules in other
languages.A package consists of one or more .go source files in a single directory that
define what the package does. Each source file begins with a package declaration, here
package main
*/
package main

/*
Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
The format 'verbs' are derived from C's but are simpler.
*/

import "fmt"

func main() {
	fmt.Println("hello to world of Go Language!")
	fmt.Println("You know that 1+1 is:", 1+1)
}
