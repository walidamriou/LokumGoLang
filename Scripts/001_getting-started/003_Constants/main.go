/*
Lokum GoLang
Learn Go Programming Language with small programs examples
----------
by Walid Amriou
https://lokumgolang.walidamriou.com
----------
How to use:
- To run the program use go run:
$ go run main.go
The Pi number is 3.14
You know? the world word in chinese is: 世界

- To build it to binaries. use go build:
$ go build main.go
$ ls
main  main.go

- To execute the built binary after build it use ./helloworld:
$ ./main
The Pi number is 3.14
You know? the world word in chinese is: 世界
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

const Pi = 3.14

func main() {
	fmt.Println("The Pi number is", Pi)

	const World = "世界"

	fmt.Println("You are know? the world word in chinese is:", World)
}
