// Every go file needs to belong to a package
// Let's name our package 'main'
package main

// This is an existing package containing functions
// that we want to use
import "fmt"

// Reusable, functional unit of code
// Serves as an entry point to the package
// No input parameters
func main() {
	// fmt - format - this is anothe rpackage of the standard library of Go
	// Println - it's a function of the fmt package
	//         - it writes to the standard output, followed by a newline
	fmt.Println("Hello Followers!")
}
