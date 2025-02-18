package main

import "fmt"

var globalVarCity string = "Agra"
var globalCityCount int

func main() {

	fmt.Println("This is Variable example!!")

	fmt.Printf("Global Variable globalVarCity = %v and globalCityCount = %v \n", globalVarCity, globalCityCount)

	//In Golang, there are several ways to declare variables.
	// 1. Using 'var' keyword - Used for package-level or function-level variable declarations.
	var x = 23 // No need to declare Data type

	var count int // Defaults to 0, No need to assign Value

	var a, b, c int // Multiple variable declarations

	var z, y = 10, "hello" // Type inferred - the compiler automatically determines the type of a variable based on the value assigned to it.
	//  Above exaple equivalent to var x int = 10 and var y string = "hello"

	var personName = "Ajay" // declaring string variable

	var isValid = true //declaring boolean

	fmt.Printf("Value of variable x = %v , count = %v , a = %v , b = %v , c = %v, z = %v, y = %v , perosnName = %v and isValid = %v\n", x, count, a, b, c, z, y, personName, isValid)

	// 2. Short Declaration (:=) - This is a shorthand for declaring and initializing a variable inside a function.
	cityName := "Agra" // No need for var keyword
	cityCount := 10    //  The type is inferred automatically.
	isReady := false

	fmt.Printf("Value of cityName = %v, cityCount = %v and isReady = %v", cityName, cityCount, isReady)

}
