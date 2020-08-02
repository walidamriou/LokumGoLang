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

// define four variables type int and with values
// varB=10 ; varC=20 ; varD=30 ; varE=40
var varB, varC, varD, varE = 10, 20, 30, 40

// shortcut declaration
//varAA = varA = 15
// you can't define varAA like that: varAA := varA outside the body of main()
var varAA = varA

// shortcut declaration , multi variables
///varBB = varB = 10 ; varCC=60
var varBB, varCC = varB, 60

// define by group
var (
	time = 1465
	id   = 566
	name = "olfa"
)

func main() {
	// Any variable define here is a local variable

	// You can't define a local variable and not use it.
	// If you define a local variable without using it the compiler returns "varname declared and not used"

	// define the value of var0
	var0 = 100

	// change varAA value from 15 to 500
	varAA = 500

	// define four variables without type
	// var11, var22, var33, var44 := 1, 2, 1, 5
	// but you can't run the last line without error because you can't declared variable here
	// without use it

	/*
	 Example of blank variable, here the divide return quotient and remainder
	 but we don't need the quotient so we use _ to ignore it
	 In other lang you need to declared a variable to recieve the quotient and declared a
	 variable mean you devlared a memory space. you don't need here quatient so you don't
	 need to save memory space for it, so just use blank variable to recieve it. */
	//_, remainder := divide(15, 4)
	// to use divide you need to define it, but here just to know what mean of blank variable

	// here we create a variable without name and ignore it :D
	_ = varAA

	// We will print the values of the four variables varB, varC, varD, varE
	fmt.Printf("varB=%d ,varC=%d ,varD=%d ,varE=%d \n", varB, varC, varD, varE)

}
