// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import "fmt"

func main () {
	// Here goes your code
	names := [5]string{"Allen", "Jaspter"}
	for _, v := range names {
		fmt.Println(v)
}
}
