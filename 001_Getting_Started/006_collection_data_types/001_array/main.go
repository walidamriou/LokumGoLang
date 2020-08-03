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

var arr1 [10]int // create a array with 10 elements, note that you can't change the length of the array
// you can't do this here: arr[0]=25 ; you need to do it in the main function only

func main() {
	arr1[0] = 25
	arr1[9] = 63
	fmt.Println("The first element of the array 1 is: ", arr1[0])
	fmt.Println("The last element of the array 1 is: ", arr1[9])

	// other method to defin array
	arr2 := [4]int{15, 52, 36} // if you didn't assigned all the valures it is get 0 by the default
	fmt.Println("The first element of the array 2 is: ", arr2[0])
	fmt.Println("The last element of the array 2 is: ", arr2[3])

	// defin array without length
	arr3 := [...]int{12, 22, 36}
	fmt.Println("The first element of the array 3 is: ", arr3[0])
	fmt.Println("The last element of the array 3 is: ", arr3[2])
}
