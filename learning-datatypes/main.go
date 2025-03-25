package main

import "fmt"

// 1. Basic data types
func basicDataType() {
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

	fmt.Printf("\n Boolean Data Type isValid =%v , \n Integer Data type count =%v , \n Uint Data type totalTickets = %v , \n Float Data type pi=%v , \n Complex Data tpe com=%v", isValid, count, totalTickets, pi, com)

}

// 2. Composite Data Types - Array (fixed size)
func arrayDataType() {

	// 2.a Array: Fixed-size collection of elements, e.g., [5]int
	var cities [5]string // An String array of size 5 (default values are 0)

	cities[0] = "Agra"
	cities[3] = "Pune"
	fmt.Printf(" Arrays Data type Cities value = %v", cities)

	// Other ways to declare array
	// You can initialize an array while declaring it.
	// var citis = [3]string{"PUNE", "AGRA", "DELHI"}

	// Short Declaration Using :=
	// arr := [3]string{"Go", "Python", "Java"}

	// Using ... to Let Compiler Infer Size
	arr := [...]int{10, 20, 30, 40}
	fmt.Printf("\n arr value = %v", arr)

	// Declaring and Initializing a Multi-dimensional Array
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// matrix := [...][3]int{{7, 8, 9}, {10, 11, 12}}

	fmt.Println(matrix)

	// Notes
	// Arrays in Go have a fixed size and cannot be resized.
	// Arrays are value types, meaning they are copied when assigned to another variable.

}

func main() {

	fmt.Println("Learning Different Data types in Golang!!!")
	fmt.Println("1. Basic Data type: ")
	basicDataType()

	fmt.Println("\n\n2. Composite Data type: ")
	arrayDataType()

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

// Q: Go lang is statically typed language?
// That means we need to tell go compiler , the data type when declaring the variables
// That is either asign variable some value or its data type
//    1. var cityName = "Pune"  // Go can type inferred - automatically assign type based on its value
// or 2. var cityName string    // or assign its data type
