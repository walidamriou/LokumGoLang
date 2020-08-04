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

// We can say that Slice is a the dynamic array because there isn't dynamic array in array
// but it isn't really a dynamic array, it's a reference type; slice hasn't length

package main

import (
	"fmt"
)

var slice1 []int
var slice2 []byte

func main() {
	slice1 = []int{1, 2, 1, 56, 32, 5, 17, 58, 69, 2, 1, 45, 63}
	fmt.Printf("The first element of the Slice 1 is: %d \n", slice1[0])
	fmt.Printf("The element 6 of the Slice 1 is: %d \n", slice1[5])

	slice2 = []byte{'A', 'B', 'Z', 'T'}
	fmt.Printf("The first element of the Slice 2 is: %c \n", slice2[0])
	fmt.Printf("The element 4 of the Slice 2 is: %c \n", slice2[3])

	arr1 := [3]int{1, 65, 41}
	var slice3 []int
	slice3 = arr1[0:2] // we can use aar1[:2] the opposite is correct also, and we can get all array by arr1[:]
	fmt.Printf("The first element of the Slice 3 is: %d \n", slice3[0])
	fmt.Printf("The element 2 of the Slice 3 is: %d \n", slice3[1])

}
