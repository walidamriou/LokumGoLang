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

import (
	"fmt"
	"reflect"
)

// boolean
var a1 bool = true

// Integers
var a2 int     // signed integer (32 or 64 bit depends on operating system)
var a3 uint    // unsigned integer (32 or 64 bit depends on operating system)
var a4 rune    // The rune type is an alias for int32 , and is used represents the ASCII 128 characters, by identified by the code points 0–127.
var a5 int8    // -128 to 127
var a6 int16   // −32,768 to 32,767
var a7 int32   // −2,147,483,648 to 2,147,483,647
var a8 int64   // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
var a9 byte    // the byte type is an alias of unit8 ; 0 to 255
var a10 uint8  // 0 to 255
var a11 uint16 // 0 to 65,535
var a12 uint32 // 0 to 4,294,967,295
var a13 uint64 // 0 to 18,446,744,073,709,551,615

// this operation is make error because the types of the valures is deffirenet
// var a = a2 + a3

// Fractions (float)
// there isn't type call float
var a14 float32
var a15 float64 // default when you use brief statement

// Complex numbers
// Complex number = Real + Imaginary
var a18 complex128 // the default and it has 64 bit real and 64 bit imaginary
var a19 complex64  // 32 bit real and 32 bit imaginar

// String
var a20 string

// Error
// This type use it to handle our errors or ignore them
// If you are comming from C/C++ you know how do you will use it
var a21 error

func main() {
	// boolean
	fmt.Println("the a1 is: ", a1, " and It's type is: ", reflect.TypeOf(a1))
	// boolean has only value of "true" or "false" so we change a1 to "false" from "true"
	a1 = false
	fmt.Println("the a1 is: ", a1, " and It's type is: ", reflect.TypeOf(a1))

}
