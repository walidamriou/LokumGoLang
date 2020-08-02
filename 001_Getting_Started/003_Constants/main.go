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

// define pi with float32 type
const pi float32 = 3.1415926

// define 2020 without type
const year = 2020

func main() {
	// we use %f to print float value
	fmt.Printf("The pi is =%f \n", pi)

	//we know that 2020 is int type
	fmt.Printf("The year is =%d \n", year)

}
