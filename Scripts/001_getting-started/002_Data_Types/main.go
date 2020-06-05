/*
Lokum GoLang
Learn Go Programming Language with small programs examples
----------
by Walid Amriou
https://lokumgolang.walidamriou.com
----------
How to use:
- To run the program use go run:
$ go run data_types.go
bool var1: bool, 1
uint8 var2: uint8, 1
uint16 var3: uint16, 2
uint32 var4: uint32, 4
uint64 var5: uint64, 8
int8 var6: int8, 1
int16 var7: int16, 2
int32 var8: int32, 4
int64 var9: int64, 8
float32 var10: float32, 4
float64 var11: float64, 8
complex64 var12: complex64, 8
complex128 var13: complex128, 16
byte var14: uint8, 1
rune var15: int32, 4
uint var16: uint, 8
int var17: int, 8
uintptr var18: uintptr, 8
string var19: string, 16
string var20: string, 16
string var21: string, 16

- To build it to binaries. use go build:
$ go build data_types.go
$ ls
data_types  data_types.go

- To execute the built binary after build it use ./helloworld:
$ ./data_types
bool var1: bool, 1
uint8 var2: uint8, 1
uint16 var3: uint16, 2
uint32 var4: uint32, 4
uint64 var5: uint64, 8
int8 var6: int8, 1
int16 var7: int16, 2
int32 var8: int32, 4
int64 var9: int64, 8
float32 var10: float32, 4
float64 var11: float64, 8
complex64 var12: complex64, 8
complex128 var13: complex128, 16
byte var14: uint8, 1
rune var15: int32, 4
uint var16: uint, 8
int var17: int, 8
uintptr var18: uintptr, 8
string var19: string, 16
string var20: string, 16
string var21: string, 16
----------
The types in Go is:
- Boolean types
bool (true or false, default value false)
- Numeric types
  -- Integer Types:
  uint8    Unsigned 8-bit integers (0 to 255)
  uint16   Unsigned 16-bit integers (0 to 65535)
  uint32   Unsigned 32-bit integers (0 to 4294967295)
  uint64   Unsigned 64-bit integers (0 to 18446744073709551615)
  int8     Signed 8-bit integers (-128 to 127)
  int16    Signed 16-bit integers (-32768 to 32767)
  int32    Signed 32-bit integers (-2147483648 to 2147483647)
  int64    Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

  -- Floating Types:
  float32    IEEE-754 32-bit floating-point numbers
  float64    IEEE-754 64-bit floating-point numbers
  complex64  Complex numbers with float32 real and imaginary parts
  complex128 Complex numbers with float64 real and imaginary parts

  -- Others types:
  byte     same as uint8
  rune     same as int32
  uint     32 or 64 bits
  int      same size as uint
  uintptr  an unsigned integer to store the uninterpreted bits of a pointer value


- String types
string  (default value is "")

----------
Go programmers call verbs in Printf:
%d     decimal integer
%x, %o, %b     integer in hexade cimal, octal, binar y
%f, %g, %e     floating-p oint number: 3.141593 3.141592653589793 3.141593e+00
%t     boole an: true or false
%c     rune (Unico de co de point)
%s    string
%q     quoted string "abc" or rune 'c'
%v     any value in a natural for mat
%T     type of any value
%%     literal percent sig n (no operand)
----------
/*

/*
Go code is organize d into packages, which are similar to libraries or modules in other
languages.A package consists of one or more .go source files in a single directory that
define what the package does. Each source file begins with a package declaration, here
package main
*/
package main

/*
Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
The format 'verbs' are derived from C's but are simpler.
*/

import (
	"fmt"
	"unsafe"
)

func main() {
	var0 := 50 //var wihout type
	fmt.Printf("var0: %T, %d\n", var0, unsafe.Sizeof(var0))

	var var1 bool
	fmt.Printf("bool var1: %T, %d\n", var1, unsafe.Sizeof(var1))

	var var2 uint8
	fmt.Printf("uint8 var2: %T, %d\n", var2, unsafe.Sizeof(var2))

	var var3 uint16
	fmt.Printf("uint16 var3: %T, %d\n", var3, unsafe.Sizeof(var3))

	var var4 uint32
	fmt.Printf("uint32 var4: %T, %d\n", var4, unsafe.Sizeof(var4))

	var var5 uint64
	fmt.Printf("uint64 var5: %T, %d\n", var5, unsafe.Sizeof(var5))

	var var6 int8
	fmt.Printf("int8 var6: %T, %d\n", var6, unsafe.Sizeof(var6))

	var var7 int16
	fmt.Printf("int16 var7: %T, %d\n", var7, unsafe.Sizeof(var7))

	var var8 int32
	fmt.Printf("int32 var8: %T, %d\n", var8, unsafe.Sizeof(var8))

	var var9 int64
	fmt.Printf("int64 var9: %T, %d\n", var9, unsafe.Sizeof(var9))

	var var10 float32
	fmt.Printf("float32 var10: %T, %d\n", var10, unsafe.Sizeof(var10))

	var var11 float64
	fmt.Printf("float64 var11: %T, %d\n", var11, unsafe.Sizeof(var11))

	var var12 complex64
	fmt.Printf("complex64 var12: %T, %d\n", var12, unsafe.Sizeof(var12))

	var var13 complex128
	fmt.Printf("complex128 var13: %T, %d\n", var13, unsafe.Sizeof(var13))

	var var14 byte
	fmt.Printf("byte var14: %T, %d\n", var14, unsafe.Sizeof(var14))

	var var15 rune
	fmt.Printf("rune var15: %T, %d\n", var15, unsafe.Sizeof(var15))

	var var16 uint
	fmt.Printf("uint var16: %T, %d\n", var16, unsafe.Sizeof(var16))

	var var17 int
	fmt.Printf("int var17: %T, %d\n", var17, unsafe.Sizeof(var17))

	var var18 uintptr
	fmt.Printf("uintptr var18: %T, %d\n", var18, unsafe.Sizeof(var18))

	var var19 string
	fmt.Printf("string var19: %T, %d\n", var19, unsafe.Sizeof(var19))

	var var20, var21 string
	fmt.Printf("string var20: %T, %d\n", var20, unsafe.Sizeof(var20))
	fmt.Printf("string var21: %T, %d\n", var21, unsafe.Sizeof(var21))
}
