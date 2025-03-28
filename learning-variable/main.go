package main

import "fmt"

var globalVarCity string = "Agra" //  In Go, global variables are declared outside functions but within the package scope.
var globalCityCount int           //  These variables can be accessed from any function within the same package.
var gName = "city"

// Few point of Global variable -
// Declared outside any function.
// Accessible throughout the package.
// Can be modified from any function within the package.

// init() is a special function that initializes global variables before main().
func init() {
	globalCityCount = 100 // Initialize global variable
}

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

	//-----------------------------------------------------------------------------------------------------------------
	// 2. Short Declaration (:=) - This is a shorthand for declaring and initializing a variable inside a function.
	cityName := "Agra" // No need for var keyword
	cityCount := 10    //  The type is inferred automatically.
	isReady := false
	// Note: With Short Declaration , while declaratio you can have to provide some value otherwise it will through error
	// e.g cityCount := int - error , cityCount := 20 - success
	// Most of cases, Short declaration generally used when function return some value
	// e.g areaCircle := areaOfCircle(12, 23) - you dont need to declare var areaCircle seperatly

	fmt.Printf("Value of cityName = %v, cityCount = %v and isReady = %v", cityName, cityCount, isReady)

	//-----------------------------------------------------------------------------------------------------------------
	// 3. Declaring Multiple Variables at Once (Grouped Declaration)
	//   This is useful for organizing related variables.
	var (
		age     int
		empName string  = "John"
		salary  float64 = 45000.50
	)
	// With short declaration =   age, empName, salary := 1, "John", 45000.50
	fmt.Println("age =", age, " , Employee Name= ", empName, " and salary= ", salary)

	//-----------------------------------------------------------------------------------------------------------------
	// 4. Constants with (const) keyword
	const pi = 3.14
	const name string = "Golang"
	// Use const for immutable values.

	//-----------------------------------------------------------------------------------------------------------------
	// 5. Blank Identifier (_)
	// Used when a variable is declared but not needed.
	i, _ := randamNumber() // Ignore the second return value
	fmt.Println("Random number i=", i)

	//-----------------------------------------------------------------------------------------------------------------
	// 6. Using Pointers
	var ptr *int // Pointer to an int
	ptr = &x     // Assign address of x

	fmt.Printf("Pointer value = %v", ptr)

	//Note : You will use proper pointer variable, when you use 'struct'
}

func randamNumber() (int, int) {
	return 12, 45
}

// Few Imp Question -
// 1. **why Does Go's Compiler Allow Unused const, But Not Unused Variables?**
// Go enforces variable usage because var occupies memory at runtime, and unused variables may indicate bugs or unnecessary code.
// In contrast, const values are evaluated at compile-time and do not take up memory at runtime, so the compiler allows them to be
// unused without affecting performance or execution.

// 2. Can We Use Short Declaration (:=) for Global Variables in Golang?
// No, we cannot use := for global variables in Go. It is only allowed inside functions.
// := is shorthand for var and only works inside functions.
// It is designed for local scope (inside func), not at the package level.
// If := were allowed for global variables, it could create ambiguity in scope.
// Local functions might unintentionally shadow global variables, making it unclear whether a variable is being modified or redefined.
// Go enforces var for global declarations to ensure explicit scoping and avoid unintended shadowing.
