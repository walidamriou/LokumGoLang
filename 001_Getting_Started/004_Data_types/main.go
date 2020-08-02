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
// Go uses by default UTF-8
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

	// Integers
	a12 = 4294967295
	fmt.Println("The a14 is: ", a12, " and It's type is: ", reflect.TypeOf(a12))

	// Fractions (float)
	a14 = 0.25
	fmt.Println("The a14 is: ", a14, " and It's type is: ", reflect.TypeOf(a14))

	// Complex numbers
	a18 = 16 + 25i
	fmt.Println("The a18 is: ", a18, " and It's type is: ", reflect.TypeOf(a18))

	// String
	// Go uses by default UTF-8
	a20 = "Hi, I'm Walid :D "
	fmt.Println("Text: ", a20, " and It's type is: ", reflect.TypeOf(a20))

	// change a character from string
	// you can't change a character from the string object and if you want to change it you will need this operation
	a20_edit := []byte(a20)     // convert a10 to []byte type
	a20_edit[0] = 'G'           // note that a20_edit is just a name :)
	a20_new := string(a20_edit) // Covert back to string type but with new valure
	fmt.Println("Text: ", a20_new, " and It's type is: ", reflect.TypeOf(a20_new))

	// add string to string
	a20_add := a20 + a20_new
	fmt.Println("Text: ", a20_add, " and It's type is: ", reflect.TypeOf(a20_add))

	// Note: Go lang will tell you "don't use underscores in Go names" but for me I like
	// this style of names, and I used it in all my codes :) so change it to what Go lang say
	// but for me, I couldn't change it because it's my coding style :p

	// substring
	a20_sub := a20[:8] + "olfa" // returns substring from index 0 to the index 8 with "olfa"
	fmt.Println("Text: ", a20_sub, " and It's type is: ", reflect.TypeOf(a20_sub))

	// multilines with string
	text := `Hi mr developer you are learn
	Golang`
	fmt.Println("Text: ", text, " and It's type is: ", reflect.TypeOf(text))

}
