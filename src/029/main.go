package main

import "fmt"

var y = 42
var z = "Shaken"

// Static programming language
// not a dynamic programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	//z = 43 Cannot assing int to string.
}
