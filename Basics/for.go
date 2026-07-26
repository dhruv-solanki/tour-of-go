package main

// The basic for loop has three components separated by semicolons:

// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum", sum)

	// the init and post statements are optional
	// so for loop is basically now while loop
	// if you omit the condition it will loop forever
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println("Sum", sum)
}
