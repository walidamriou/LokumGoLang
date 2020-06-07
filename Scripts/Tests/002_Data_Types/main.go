/*
Lokum GoLang
Learn Go Programming Language with small programs examples
----------
by Walid Amriou
https://lokumgolang.walidamriou.com
----------
*/

package main

// any variable define here is global variable

// define with "var0" as name and int as type
// var var0 int

// define variable with name "varA", type int and value 15
// varA = 15
// var varA int = 15

// define four variables type int
// var var1, var2, var3, var4 int

// define four variables without type
// var var11, var22, var33, var44 int

// define four variables type int and with valures
// varB=10 ; varC=20 ; varD=30 ; varE=40
// var varB, varC, varD, varE = 10, 20, 30, 40

// shortcut declaration
//varAA = varA = 15
// varAA := varA

// shortcut declaration , multi variables
// varBB = varB = 10 ; varCC=60
// varBB, varCC := varB, 60

func main() {
	// any variable define here is local variable

	// define the value of var0
	// var0 = 100

	// change varAA valure from 15 to 500
	// varAA = 500

	// example of black variable, here the divide return quotient and remainder
	// but we don't need the quotient so we use _ to ignore it
	// _ , remainder := divide(15,4)

	//fmt.Printf("var0: %T, %d\n", var0, unsafe.Sizeof(var0))

}
