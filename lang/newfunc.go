// newfunc
// alternative way to create variables
package main

import (
	"fmt"
)

func main() {
	// new inits a variable with zero value, and returns a pointer to the variable
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	// each new() returns a distinct address if there are zero values of the type
	a := new(int)
	b := new(int)
	fmt.Println(a == b) // false

	// exceptions: no data, no zero value, so same address
	c := new(struct{})
	d := new(struct{})
	fmt.Println(c == d) // true
	e := new([0]int)    // 0-length int array
	f := new([0]int)
	fmt.Println(e == f)
}

// identical to newFunc2
func newFunc1() *int {
	return new(int)
}

// identical to newFunc1
func newFunc2() *int {
	var dummy int
	return &dummy
}
