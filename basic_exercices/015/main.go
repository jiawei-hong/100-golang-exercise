// Exercise: User input
// Get a number from the console and check if it's odd
// You can create another function or do it everything in the main func :)

package main

import "fmt"

func main() {
	// Here goes your code
	var number int
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("is even")
	} else {
		fmt.Println("is odd")
	}
}
