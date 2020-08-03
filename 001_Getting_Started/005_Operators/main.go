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

func main() {

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

	/*
		%	remainder	integers
		&	bitwise AND
		|	bitwise OR
		^	bitwise XOR
		&^	bit clear (AND NOT)
		<<	left shift	integer << unsigned integer
		>>	right shift	integer >> unsigned integer
	*/

}
