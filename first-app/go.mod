// When writing a Golang application, the first step is to initialize the Go module. 
// This can be done by running the following command:
// --> go mod init <your-app-name>

// This command creates a new module, which helps manage dependencies in a structured manner, similar to the pom.xml file in a Java application.
// The go.mod file is where Go stores information about dependencies. It also tracks the versions of the dependencies used in the project. 
// The go.mod file includes details such as the project name, Go version, and dependency information.

module first-app

go 1.19
