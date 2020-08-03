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

var a = 10
var b = 20
var c = 30

var e = byte(0b000110111)
var f = byte(0b010010101)

func main() {

	// Arithmetic operators
	// sum	integers, floats, complex values, strings
	d1 := a + b
	fmt.Println("a + b = ", d1)

	// difference	integers, floats, complex values
	d2 := a - b
	fmt.Println("a - b = ", d2)

	// product
	d3 := a * b
	fmt.Println("a * b = ", d3)

	// quotient
	d4 := b / a
	fmt.Println("b / a = ", d4)

	// remainder
	d5 := a % b
	fmt.Println("a % b = ", d5)

	// Logical operators
	// AND
	d6 := 1 & 0
	fmt.Println("1 AND 0 = ", d6)

	// OR
	d7 := 1 | 0
	fmt.Println("1 OR 0 = ", d7)

	// bitwise XOR
	d8 := 1 ^ 0
	fmt.Println("1 XOR 0 = ", d8)

	// AND NOT
	d9 := 1 &^ 0
	fmt.Println("1 AND NOT 0 = ", d9)

	// left shift
	d10 := 1 << 0
	fmt.Println("1 left shift 0 = ", d10)

	// right shift
	d11 := 1 >> 0
	fmt.Println("1 right shift 0 = ", d11)

}
