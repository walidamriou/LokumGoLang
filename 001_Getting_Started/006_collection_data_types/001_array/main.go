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
	arr3 := [...]int{12, 22, 36} // Golang calculates the length and you can't change it after that (this is not a dynamic)
	fmt.Println("The first element of the array 3 is: ", arr3[0])
	fmt.Println("The last element of the array 3 is: ", arr3[2])

	// array multi-dimensional
	/*
		example of two dimensionals
		   0  1
		0 12 54
		1 25 65
		2 85 20

		[0][0] is 12
		[2][1] is 20
	*/
	arr4 := [3][2]int{[2]int{12, 54}, [2]int{25, 65}, [2]int{85, 20}}
	fmt.Println("The first element of the array 4 is: ", arr4[0][0])
	fmt.Println("The last element of the array 4 is: ", arr4[2][1])

	// other method
	arr5 := [3][2]int{{12, 54}, {25, 65}, {85, 20}}
	fmt.Println("The first element of the array 5 is: ", arr5[0][0])
	fmt.Println("The last element of the array 5 is: ", arr5[2][1])
}
