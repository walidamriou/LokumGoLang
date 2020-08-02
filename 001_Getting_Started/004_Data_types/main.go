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
// basic
var a2 int  // signed integer (32 or 64 bit depends on operating system)
var a3 uint // unsigned integer (32 or 64 bit depends on operating system)
var a4 rune // The rune type is an alias for int32 , and is used represents the ASCII 128 characters, by identified by the code points 0–127.

func main() {

	fmt.Println("the a1 is: ", a1, " and It's type is: ", reflect.TypeOf(a1))
	// boolean has only value of "true" or "false" so we change a1 to "false" from "true"
	a1 = false
	fmt.Println("the a1 is: ", a1, " and It's type is: ", reflect.TypeOf(a1))

}
