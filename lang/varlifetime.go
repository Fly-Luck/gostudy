// varlifetime
// life is short for variables in stack, but longer for those in heap
package main

var global *int

/*
func main() {
	f()
	g()
}
*/

func f() {
	// x is allocated in heap, because it's reachable after f is finished
	// x escapes from f
	var x int
	x = 1
	global = &x
}

func g() {
	// y is allocated in stack, and may be recycled after g is finished
	// y is not reachable even though it's created by new()
	y := new(int)
	*y = 1
}
