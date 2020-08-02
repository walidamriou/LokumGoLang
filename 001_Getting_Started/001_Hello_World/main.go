/*
 ******************************************************************** 
  Lokum GoLang
  To Learn Go Programming Language with small programs examples

  Project Website: lokumgolang.walidamriou.com
  Github: https://github.com/walidamriou/LokumGoLang
  Copyright CC 2020 Walid Amriou
  This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License
   
  Last update: August 2020
 ******************************************************************** 
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
	// Print hello to world of Go Language! in Console
	fmt.Println("hello to world of Go Language!")
	// Print the result of 1+1
	fmt.Println("You know that 1+1 is:", 1+1)
}
