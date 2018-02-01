// assignments
// tuple assignments
package main

import (
	"fmt"
)

func main() {
	//x = 1 // named variable
	//*p = true //indirect variable
	//person.name = "bob" // struct field
	//count[x] *= count[x] * scale // array/slice/map element

	// tuple assignments
	// basic
	v1, v2, v3 := 2, 3, 5
	fmt.Println(v1, v2, v3)

	// swap x y
	x := 2
	y := 6
	x, y = y, x
	fmt.Println(x, y)

	// greates common divisor
	a, b := x, y
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a)

	// n-th Fibonacci number iteratively
	n := 20
	c, d := 0, 1
	for i := 0; i < n; i++ {
		c, d = d, c+d
	}
	fmt.Println(c)

	// typical error handling
	//f, err = os.Open("foo.txt")

	// operators that return 2 results
	//v, ok = m[key] //map lookup
	//v, ok = x.(T) //type assertion
	//v, ok = <-ch // channel receive

	// ignoring unwanted values
	//_, err = io.Copy(dst, src)
	//_, ok = x.(T)
}
