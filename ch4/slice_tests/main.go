// This package is for testing the rev package imported below
package main

import (
	"The-Go-Programming-Language/ch4/rev"
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	// Calling reverse once just reverses
	rev.Reverse(a[:])
	fmt.Println(a)

	// However, calling it as below "rotates" the
	// slice by several positions
	rev.Reverse(a[:2])
	rev.Reverse(a[2:])
	rev.Reverse(a[:])
	fmt.Println(a)
}
