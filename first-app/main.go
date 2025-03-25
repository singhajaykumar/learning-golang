// Writing a Go application starts with initializing a Go module.
// This can be done by running the following command:
// --> go mod init <your-app-name>
//
// This command creates a new module, which helps manage dependencies in a structured manner,
// similar to the pom.xml file in a Java application.
// The go.mod file stores information about dependencies and tracks their versions used in the project.
// It includes details such as the project name, Go version, and dependency information along with their version.

// In Go, all code is organized into packages.
// The first line specifies the package name.
package main

// To import a package, we use the 'import' keyword followed by the package name.
// The fmt package is used for formatted input and output operations.
import "fmt"

// The main function serves as the entry point of a Go application.
// Every Go application must have a main function inside the 'main' package.
func main() {
	fmt.Println("Hello, World!") // Prints "Hello, World!" to the console.
}

// To execute a Go application, use the following command:
// --> go run <application-name>

// Note: In Go, a package is a way to organize and reuse code. Every Go program is built using packages.
// Library Packages: Any package other than 'main' is considered a library package.
