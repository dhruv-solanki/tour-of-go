package main

import "fmt"

var golang int

// A var declaration can include initializers, one per variable.

// If an initializer is present, the type can be omitted;
// the variable will take the type of the initializer.
var c, python, java = true, false, "no!"

func main() {
	var i, j = 1, 2

	// Inside a function, the := short assignment statement
	// can be used in place of a var declaration with implicit type.

	// Outside a function, every statement begins with a keyword
	// (var, func, and so on) and so the := construct is not available.
	k := 3
	greet := "Hello World!"
	fmt.Println(greet)

	fmt.Println(i, j, k, c, python, java, golang)
}
