package main

import "fmt"

func main() {

	fmt.Println("Learning Different Data types in Golang!!!")

	// There are 3 basic data types
	// 1. Boolean: bool (true or false)
	var isValid bool = true

	// 2. Numeric Types: it further have 4 sub-types
	// 2.a integer: int, int8, int16, int32, int64
	var count int

	// 2.b Unsigned Integer: uint, uint8, uint16, uint32, uint64, uintptr
	var totalTickets uint = 50 // can't be negative number
	//totalTickets = -20  //through error

	// 2.c Floating Point: float32, float64
	var pi float32 = 3.14

	// 2.d Complex Numbers: complex64, complex128
	//  complex data type represents complex numbers, which consist of a real and an imaginary part, both of type float32.
	var com complex64 = 5 + 0i

	fmt.Printf("isValid =%v , count =%v , totalTickets = %v , pi=%v , and com=%v/n", isValid, count, totalTickets, pi, com)

	// Special Data Types
	// nil: Represents an uninitialized reference type (e.g., pointer, slice, map, interface, channel, function).
	// rune: Alias for int32, used for Unicode characters.
	// byte: Alias for uint8, used for raw byte storage.

}

// Notes:
// the default data type for int is platform-dependent:
// On a 32-bit system, int is equivalent to int32 (4 bytes).
// On a 64-bit system, int is equivalent to int64 (8 bytes).
// If you want to print whether int is treated as int32 or int64
// based on your machine's platform, you can use the unsafe package to check the size of int
// var count int
// fmt.Printf("Type of x: int%v (size: %d bytes)\n", unsafe.Sizeof(count)*8, unsafe.Sizeof(count))
