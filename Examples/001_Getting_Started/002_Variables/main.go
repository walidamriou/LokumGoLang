/*
Lokum GoLang
Learn Go Programming Language with small programs examples
----------
by Walid Amriou
https://lokumgolang.walidamriou.com
----------
*/

package main

import "fmt"

// any variable define here is global variable

// define "var0" with int type
var var0 int

// define variable with name "varA", type int and value 15
// varA = 15
var varA int = 15

// define four variables type int
var var1, var2, var3, var4 int

// define four variables without type
var var11, var22, var33, var44 int

// define four variables type int and with valures
// varB=10 ; varC=20 ; varD=30 ; varE=40
var varB, varC, varD, varE = 10, 20, 30, 40

// shortcut declaration
//varAA = varA = 15
// you can't define varAA like that: varAA := varA outside the body of main()
var varAA = varA

// shortcut declaration , multi variables
///varBB = varB = 10 ; varCC=60
var varBB, varCC = varB, 60

func main() {
	// any variable define here is local variable
	// You can't define a local variable without use it.
	// If define a local variable without use it the compiler return "varname declared and not used"

	// define the value of var0
	var0 = 100

	// change varAA valure from 15 to 500
	varAA = 500

	// example of blank variable, here the divide return quotient and remainder
	// but we don't need the quotient so we use _ to ignore it
	//_, remainder := divide(15, 4)
	// to use divide you need to define it, but here just to know what mean of blank variable
	// here we create a variable without name and ignore it :D
	_ = varAA

	// We will print the valures of the four variables varB, varC, varD, varE
	fmt.Printf("varB=%d ,varC=%d ,varD=%d ,varE=%d \n", varB, varC, varD, varE)

}
