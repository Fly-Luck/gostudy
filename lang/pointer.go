// pointer
// basic pointer usages
package main

/*
import (
	"fmt"
)
*/

/*
func main() {
	a := 1          // int variable
	p := &a         // pointer to the variable
	fmt.Println(*p) // value of the variable to which the pointer points
	*p = 2          // change the value of the variable to which the pointer points, via the pointer
	fmt.Println(a)  // new value of the variable

	var x, y int // variables of zero values
	// 1. &x returns the pointer of variable x
	// 2. zero value of a pointer is nil
	// 3. if a pointer points to a varible, it's NOT nil
	// 4. two pointers are equal only if they points to the same variable or both nil
	fmt.Println(&x == &x, &x == &y, &x == nil)

	p2v := f()
	fmt.Println(*p2v)
	// local variable in function is not the same, thus the pointers returned are different
	fmt.Println(f() == f())

	v := 1
	incr(&v)              // v is 2
	fmt.Println(incr(&v)) // v is 3
}
*/

// the value of variable is stored if a pointer is returned
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	// 1. *p +1
	// 2. p not changed
	*p++
	return *p
}
